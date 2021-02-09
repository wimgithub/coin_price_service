package test

import (
	CrossEther "coin_price_service/ether_contract"
	model "coin_price_service/models"
	"coin_price_service/models/mysql"
	"coin_price_service/pkg/conversion"
	"coin_price_service/pkg/gredis"
	"coin_price_service/pkg/http_util"
	"coin_price_service/pkg/setting"
	"context"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestTronEventsScan(t *testing.T) {
	con := conversion.New()
	var data struct {
		Data []*model.Data `json:"data"`
	}
	get, err := http_util.Get("https://api.trongrid.io/v1/contracts/TSigzQpjVJfyTuybBL3zGCbxJX3bvZVhoN/events?event_name=Receive&order_by=block_timestamp%2Casc&limit=200")
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}
	err = json.Unmarshal(get, &data)
	if err != nil {
		fmt.Println("Err: ", err)
		return
	}
	if data.Data[0].Result.Value == "" {
		fmt.Println("获取失败！")
		return
	}
	fmt.Println("data: ", data.Data)

	var rec []*model.RechargeEvents
	for _, v := range data.Data {
		rec = append(rec, &model.RechargeEvents{
			RechargeTransactionHash: v.TransactionId,
			BlockNumber:             v.BlockNumber,
			BlockTimestamp:          v.BlockTimestamp,
			Contract:                v.ContractAddress,
			UserAddress:             v.Result.Addr,
			Value:                   con.ToDecimal(v.Result.Value, 6),
			ChainType:               "Tron",
		})
	}
	err = mysql.SharedStore().AddRecTxs(rec)
	if err != nil {
		fmt.Println("MysqlErr: ", err.Error())
	}
}

func TestETHEventsScan(t *testing.T) {
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/9cff2e75a50c45f29f3afe2b56a795d8")
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}
	contractAddress := common.HexToAddress("0xA79f3c9aD484b7A8fFA014D6B83B95a5D993A311")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(CrossEther.CrossEtherABI)))
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}

	for {
		select {
		case err := <-sub.Err():
			fmt.Println("ERR: ", err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
			var m = make(map[string]interface{})
			err := contractAbi.UnpackIntoMap(m, "Receive", vLog.Data)
			if err != nil {
				fmt.Println("ERR: ", err)
				return
			}
			fmt.Println("addr: ", m["addr"].(common.Address).String())
			fmt.Println("value: ", m["value"])
		}
	}

}

func TestETHEventsBlockNumberScan(t *testing.T) {
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/9cff2e75a50c45f29f3afe2b56a795d8")
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}
	number, err := client.BlockNumber(context.Background())
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}
	fmt.Println(number)
	contractAddress := common.HexToAddress("0xA79f3c9aD484b7A8fFA014D6B83B95a5D993A311")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(7637936),
		ToBlock:   big.NewInt(int64(number)),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(CrossEther.CrossEtherABI)))
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}
	for _, v := range logs {
		var m = make(map[string]interface{})
		err := contractAbi.UnpackIntoMap(m, "Receive", v.Data)
		if err != nil {
			fmt.Println("ERR: ", err)
			return
		}
		fmt.Println("addr: ", m["addr"].(common.Address).String())
		fmt.Println("value: ", m["value"].(*big.Int).Int64())
		fmt.Println("txHash: ", v.TxHash.String())
	}
}

func TestRedis(t *testing.T) {
	setting.Setup()
	err := gredis.SharedSnapshotStore().Set(gredis.TronBlockTime, 0, 0)
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}
	_ = gredis.SharedSnapshotStore().Set(gredis.EtherBlock, 0, 0)
}

func Test(t *testing.T) {
	var coins = []string{"bsvusdt", "htusdt", "filusdt", "ethusdt", "btcusdt", "ltcusdt", "bchusdt", "dotusdt"}
	url := "https://api.huobi.pro/market/detail/merged?symbol="
	var PData model.HuoBiPrice
	ch := make(chan []byte, len(coins))
	quitChan := make(chan bool)
	for _, v := range coins {
		go func(n string) {
			fmt.Println("name: ", n)
			bytes, err := http_util.Get(url + n)
			if err != nil {
				fmt.Println(err)
				return
			}
			ch <- bytes[:]
		}(v)
	}
	quitChan <- true
	for {
		select {
		case data := <-ch:
			json.Unmarshal(data, &PData)
			fmt.Println("Close: ", PData.Tick.Close)
		case <-quitChan:
			return
		}
	}
}

func TestGetUsers(t *testing.T) {
	url := "https://api-scan.hecochain.com/hsc/listTokenHolder/0x9fe0aa47fc8ad1a255645719f653a28796bd8e18/"
	url2 := "/100?x-b3-traceid=6d394454c8fa0fb71f3306c37ad974b7"

	var d *model.HecoUsers
	i := 1
	for {
		bytes, err := http_util.Get(url + strconv.Itoa(i) + url2)
		if err != nil {
			fmt.Println(err)
			break
		}
		err = json.Unmarshal(bytes, &d)
		if err != nil || len(d.HData.Data) == 0 {
			fmt.Println(err.Error())
			return
		}
		s := []string{}
		for _, v := range d.HData.Data {
			s = append(s, v.Address)
		}
		fmt.Println(len(s))

		f, err := os.Create(strconv.Itoa(i) + "txt")
		if err != nil {
			return
		}
		marshal, _ := json.Marshal(s)
		_, err1 := io.WriteString(f, string(marshal)) //写入文件(字符串)
		if err1 != nil {
			panic(err1)
		}
		f.Close()
		i++
		time.Sleep(3 * time.Second)
	}
}

func TestGetExcel(t *testing.T) {
	xlsx, err := excelize.OpenFile("./kt.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Get value from cell by given sheet index and axis.
	//cell := xlsx.GetCellValue("Sheet1", "B2")
	//fmt.Println(cell)
	// Get sheet index.
	//index := xlsx.GetSheetIndex("Sheet1")
	// Get all the rows in a sheet.
	var a []string
	rows := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			a = append(a, colCell)
			//fmt.Print(colCell, "\t")
		}
		//fmt.Println()
	}
	marshal, _ := json.Marshal(a)
	fmt.Println(string(marshal))

}

func TestSoli(t *testing.T) {
	/*
	   for (uint256 i = 1; i < size; i++) {

	       for (uint256 j = i; j > 0 && array[j-1]  > array[j]; j--) {
	           // j = 1 ; 1 > 0 && 0.9 > 1
	           uint256 tmp = array[j];
	           array[j] = array[j-1];
	           array[j-1] = tmp;
	       }
	   }
	   if (size % 2 == 1) {
	       return array[size / 2];
	   } else {
	       return array[size / 2].add(array[size / 2 - 1]) / 2;
	   }
	*/

	a := func(a []float64, size int) {
		for i := 1; i < size; i++ {
			fmt.Println("i: ",i)
					//  2 > 0 &&  a[2-1]  > a[2]
					//  2 > 0 &&   3 > 2
			for j := i; j > 0 && a[j-1] > a[j]; j-- {
				fmt.Println("j:",j)
				tmp := a[j] //  2
				fmt.Println("tmp: ",tmp)
				a[j] = a[j-1] // 3
				a[j-1] = tmp // 2

				fmt.Println("a: ",a)
			}
		}


		fmt.Println("==============: ",a)


		if size%2 == 1 {
			fmt.Println("奇数：", a[size/2])
		} else {
			fmt.Println("偶数：", (a[size/2]+a[size/2-1])/2)
			fmt.Println("a[size/2] ",a[size/2]," + a[size/2-1] ",a[size/2-1]," / 2")
		}
	}
	a([]float64{114,110,},2)
}
