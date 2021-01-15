package scan

import (
	"context"
	"coin_price_service/coin_service/client"
	model "coin_price_service/models"
	"coin_price_service/models/mysql"
	"coin_price_service/pkg/conversion"
	"coin_price_service/pkg/gredis"
	"coin_price_service/pkg/http_util"
	"coin_price_service/pkg/logging"
	"coin_price_service/pkg/setting"
	"coin_price_service/pkg/util"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
	"strings"
	"time"
)

type TronScan struct {
	URL        string
	LastTime   int64
	AccountURL string
}

func NewTronScan() *TronScan {
	lastTime, err := gredis.SharedSnapshotStore().Get(gredis.TronBlockTime)
	if err != nil {
		lastTime = 0
	}
	return &TronScan{
		URL:        setting.TronscanSetting.Host + "contracts/" + setting.TronscanSetting.Contract + "/events?" + setting.TronscanSetting.EventName + "&" + setting.TronscanSetting.OrderBy + "&" + setting.TronscanSetting.Limit,
		AccountURL: setting.TronscanSetting.Host + "accounts/" + setting.TronscanSetting.AccountAddress + "/transactions/trc20?only_to=true&limit=" + setting.TronscanSetting.Limit + "&contract_address=" + setting.TronscanSetting.MMMContract,
		LastTime:   lastTime,
	}
}

func (t *TronScan) Start() {
	go t.account_scan()
	go t.freed_start()
	//go t.freed()
}

func (t *TronScan) account_scan() {
	con := conversion.New()
	for {
		events, err := t.getTronEvents(t.AccountURL, false)
		if err != nil {
			fmt.Println("URL: ", t.AccountURL)
			fmt.Println("ERR: ", err)
			time.Sleep(3 * time.Second)
			continue
		}
		var rec []*model.RechargeEvents
		for _, v := range events {
			if v.To != setting.TronscanSetting.AccountAddress {
				continue
			}
			rec = append(rec, &model.RechargeEvents{
				RechargeTransactionHash: v.TransactionId,
				BlockNumber:             v.BlockNumber,
				BlockTimestamp:          v.BlockTimestamp,
				Contract:                v.To,
				UserAddress:             v.From,
				Value:                   con.ToDecimal(v.Value, setting.TronscanSetting.Decimals),
				ChainType:               "Tron",
			})
		}
		err = mysql.SharedStore().AddRecTxs(rec)
		if err != nil {
			fmt.Println("MysqlErr: ", err.Error())
		}
		time.Sleep(2 * time.Second)
	}
}

func (t *TronScan) freed_start() {
	for {
		now := time.Now()
		// 计算下一个零点
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		ti := time.NewTimer(next.Sub(now))
		fmt.Println("波场兑币清算-预计下次清算时间：", next.Sub(now).String())
		select {
		case <-ti.C:
			go t.freed()
		}
	}
}

func (t *TronScan) freed() {
	record, err := mysql.SharedStore().GetFullRecord("Tron")
	if err != nil {
		logging.Error("时间: ", time.Now().Format("2006-01-02 15:04:05"), " 获取充值列表失败结束清算!")
		return
	}
	freedTotal := t.add_freed(record)
	confTotal := decimal.NewFromFloat(setting.TronscanSetting.FreedTotal)

	for _, v := range record {
		fmt.Println("FreedTotal: ", confTotal)
		if confTotal.GreaterThan(freedTotal) {
			t.freed_transfer(v.UserAddress, v.Value.Sub(v.FreedValue), v)
			fmt.Println("给 ", v.UserAddress, " 的映射地址转入 ", v.Value.Sub(v.FreedValue))
		} else {
			// 当前用户释放数量 释放总量 * (用户充值数量/充值总量)
			total := confTotal.Mul((v.Value.Sub(v.FreedValue)).Div(freedTotal))
			t.freed_transfer(v.UserAddress, total, v)
			fmt.Println("给 ", v.UserAddress, " 的映射地址转入 ", total)
		}
		time.Sleep(3 * time.Second)
	}
}

func (t *TronScan) add_freed(data []*model.RechargeEvents) (total decimal.Decimal) {
	for _, v := range data {
		total = total.Add(v.Value.Sub(v.FreedValue))
	}
	return
}

func (t *TronScan) freed_transfer(address string, value decimal.Decimal, record *model.RechargeEvents) {
	// 获取当前用户绑定的地址
	bid, _ := mysql.SharedStore().QueryBid(&model.BidAddress{TronAddress: address})
	if len(bid) == 0 {
		// or 更新状态 当前用户未绑定
		logging.Error("TronAddress: ", address, " Not Bid Ether Address！")
		return
	}
	// 转账
	status := 1
	wei := conversion.New().ToWei(value, setting.EtherscanSetting.Decimals)
	hash, key, err := t.TransferMMM(bid[0].EtherAddress, wei)
	if err != nil {
		logging.Error("TransferMMM Error: ", err)
		status = 3
	}
	// 添加资金释放记录
	_ = mysql.SharedStore().FreedRecordInsert(&model.FreedRecord{
		Address:    bid[0].EtherAddress,
		FreedValue: value,
		Hash:       hash,
		Status:     status,
		ChainType:  "Ether",
	})
	if err == nil {
		// 交易监听
		go client.EthCli.StartTransactionReceipt(hash, key, record, value)
	}
}

func (t *TronScan) scan() {
	con := conversion.New()
	url := t.URL + "&min_block_timestamp="
	fmt.Println("启动Tron 充值监听服务...")
	for {
		recs, err := t.getTronEvents(url+strconv.Itoa(int(t.LastTime)), true)
		if err != nil {
			fmt.Println("ERR: ", err)
			time.Sleep(3 * time.Second)
			continue
		}
		if len(recs) == 0 {
			time.Sleep(3 * time.Second)
			continue
		}
		var rec []*model.RechargeEvents
		for _, v := range recs {
			rec = append(rec, &model.RechargeEvents{
				RechargeTransactionHash: v.TransactionId,
				BlockNumber:             v.BlockNumber,
				BlockTimestamp:          v.BlockTimestamp,
				Contract:                v.ContractAddress,
				UserAddress:             address.HexToAddress(util.AddressDealWith(v.Result.Addr)).String(),
				Value:                   con.ToDecimal(v.Result.Value, setting.TronscanSetting.Decimals),
				ChainType:               "Tron",
			})
		}
		err = mysql.SharedStore().AddRecTxs(rec)
		if err != nil {
			fmt.Println("MysqlErr: ", err.Error())
		}
		time.Sleep(3 * time.Second)
	}
}

func (t *TronScan) getTronEvents(url string, b bool) ([]*model.Data, error) {
	var RecData = struct {
		Data []*model.Data `json:"data"`
	}{}
	get, err := http_util.Get(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(get, &RecData)
	if err != nil {
		return nil, err
	}
	if RecData.Data[0].Result.Value == "" && b == true {
		return nil, errors.New("Get Data Error!")
	}
	t.LastTime = RecData.Data[len(RecData.Data)-1].BlockTimestamp
	return RecData.Data, nil
}

// 监听快照备份
func (s *TronScan) runSnapshots() {
	for {
		select {
		case _, ok := <-gredis.RedisSnapshot:
			if !ok {
				err := gredis.SharedSnapshotStore().Set(gredis.TronBlockTime, s.LastTime, 0)
				if err != nil {
					logging.Error("tron_block_time 备份失败: ", err)
				}
				return
			}
		}
	}
}

func (t *TronScan) TransferMMM(to string, amount *big.Int) (string, string, error) {
	cli := client.EthCli
	hash := ""
	key := ""
	gasPrice, err := cli.GetEthCli().SuggestGasPrice(context.Background())
	if err != nil {
		return "", "", err
	}
	ETHWithdrawGasPrice := setting.EtherscanSetting.ETHWithdrawGasPrice
	Erc20GasLimit := setting.EtherscanSetting.Erc20GasLimit

	ETHWithdrawGasPrice = ETHWithdrawGasPrice * 100
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(int64(ETHWithdrawGasPrice))).Div(gasPrice, big.NewInt(100))

	err, hash, key = cli.TransferERC20PrivateKey(setting.EtherscanSetting.Private, gasPrice, uint64(Erc20GasLimit), amount, to, strings.ToLower(setting.EtherscanSetting.MMMContract))
	if err != nil {
		return "", "", err
	}
	return hash, key, nil
}
