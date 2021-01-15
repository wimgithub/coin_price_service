package scan

import (
	"context"
	CrossEther "coin_price_service/ether_contract"
	model "coin_price_service/models"
	"coin_price_service/models/mysql"
	"coin_price_service/pkg/conversion"
	"coin_price_service/pkg/gredis"
	"coin_price_service/pkg/logging"
	"coin_price_service/pkg/setting"
	"fmt"
	"github.com/ethereum/go-ethereum"
	abi2 "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"time"
)

type EtherScan struct {
	client    *ethclient.Client
	lastBlock int64
	address   common.Address
	abi       abi2.ABI
	eventName string
	con       *conversion.TypeConversion
}

func NewEtherScan() *EtherScan {
	block, err := gredis.SharedSnapshotStore().Get(gredis.EtherBlock)
	if err != nil {
		block = 0
	}
	client, err := ethclient.Dial(setting.EtherscanSetting.Infura)
	if err != nil {
		logging.Fatal("Ether Client Dial Err: ", err)
	}
	contractAbi, err := abi2.JSON(strings.NewReader(string(CrossEther.CrossEtherABI)))
	if err != nil {
		logging.Fatal("abi new err: ", err)
	}
	return &EtherScan{
		client:    client,
		address:   common.HexToAddress(setting.EtherscanSetting.Contract),
		lastBlock: block,
		abi:       contractAbi,
		eventName: setting.EtherscanSetting.EventName,
		con:       conversion.New(),
	}
}

func (t *EtherScan) Start() {
	fmt.Println("启动Ether 充值监听服务...")
	go t.scan()
	go t.runSnapshots()
}

func (t *EtherScan) scan() {
	t.getEvents()
	t.subEvents()
}

func (t *EtherScan) getEvents() {
	number, err := t.client.BlockNumber(context.Background())
	if err != nil {
		logging.Fatal("get block number err: ", err)
	}
	if t.lastBlock > int64(number) {
		return
	}
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(t.lastBlock),
		ToBlock:   big.NewInt(int64(number)),
		Addresses: []common.Address{
			t.address,
		},
	}
	logs, err := t.client.FilterLogs(context.Background(), query)
	if err != nil {
		logging.Fatal("filter logs err: ", err)
	}
	t.lastBlock = int64(number)
	if len(logs) == 0 {
		return
	}
	var txs []*model.RechargeEvents
	for _, v := range logs {
		var m = make(map[string]interface{})
		err := t.abi.UnpackIntoMap(m, t.eventName, v.Data)
		if err != nil {
			logging.Error("unpack into map err: ", err)
		}
		txs = append(txs, &model.RechargeEvents{
			RechargeTransactionHash: v.TxHash.String(),
			BlockNumber:             int64(v.BlockNumber),
			BlockTimestamp:          0,
			Contract:                t.address.String(),
			UserAddress:             m["addr"].(common.Address).String(),
			Value:                   t.con.ToDecimal(m["value"].(*big.Int).String(), setting.EtherscanSetting.Decimals),
			ChainType:               "Ether",
		})
	}

	err = mysql.SharedStore().AddRecTxs(txs)
	if err != nil {
		logging.Error("add rec txs err: ", err)
	}
}

func (t *EtherScan) subEvents() {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{t.address},
	}
	logs := make(chan types.Log)
	sub, err := t.client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		logging.Fatal("ether sub err: ", err)
	}
	var txs []*model.RechargeEvents
	for {
		select {
		case err := <-sub.Err():
			logging.Error("sub event err: ", err)
			continue
		case vLog := <-logs:
			var m = make(map[string]interface{})
			err := t.abi.UnpackIntoMap(m, t.eventName, vLog.Data)
			if err != nil {
				logging.Error("abi unpackIntoMap err: ", err)
			}
			t.lastBlock = int64(vLog.BlockNumber)
			txs = append(txs, &model.RechargeEvents{
				RechargeTransactionHash: vLog.TxHash.String(),
				BlockNumber:             int64(vLog.BlockNumber),
				BlockTimestamp:          0,
				Contract:                t.address.String(),
				UserAddress:             m["addr"].(common.Address).String(),
				Value:                   t.con.ToDecimal(m["value"].(*big.Int).String(), setting.EtherscanSetting.Decimals),
				ChainType:               "Ether",
			})
			fmt.Println("Value: ", t.con.ToDecimal(m["value"].(*big.Int).String(), setting.EtherscanSetting.Decimals))
			//if len(txs) > 0 && len(txs) < 2 {
			//	continue
			//}
			for {
				err := mysql.SharedStore().AddRecTxs(txs)
				if err != nil {
					logging.Error("add rec txs err: ", err)
					time.Sleep(time.Second)
					continue
				}
				txs = nil
				break
			}
		}
	}

}

func (s *EtherScan) runSnapshots() {
	for {
		select {
		case _, ok := <-gredis.RedisSnapshot:
			if !ok {
				err := gredis.SharedSnapshotStore().Set(gredis.EtherBlock, s.lastBlock, 0)
				if err != nil {
					logging.Error("ether_block 备份失败: ", err)
				}
				return
			}
		}
	}
}
