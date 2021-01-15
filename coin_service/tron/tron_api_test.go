package tron

import (
	"coin_price_service/pkg/crypto"
	"coin_price_service/pkg/util"
	"encoding/hex"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/shopspring/decimal"
	"testing"
)

// 合约查询
func TestTronCliGetContract(t *testing.T) {
	tronCli := NewTronCli()
	tronCli.Start()
	from := "TNpeo9MZW8Rc7wCf17evQMVyMGf1KdPaxK"
	contract := "TSNgPpTH2Pp7bYLct5m7TUmWrvw3JotYES"
	ofContract, err := tronCli.GetConstantResultOfContract(from, contract, util.ProcessBalanceOfParameter(from, util.GetMethodID("getUserEthByTrxAddr(address)"), false, 0))
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}
	//data := common.ToHex(ofContract[0])
	data := hex.EncodeToString(ofContract[0])
	property, err := tronCli.Torn.ParseTRC20StringProperty(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}
	fmt.Println("Data: ", property)
}

// 合约交易
func TestTronCliSendContract(t *testing.T) {
	tronCli := NewTronCli("35.180.51.163:50051")
	err := tronCli.Torn.Start()
	if err != nil {
		fmt.Println("ERR: ", err.Error())
		return
	}
	// 私钥处理
	pri, err := crypto.GetPrivateKeyByHexString("400e8ea5dbc23cf4aa03c138c7fcefcff64f68cd2e544694602136c8ea65ac10")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	// 合约地址
	addr := "TNpeo9MZW8Rc7wCf17evQMVyMGf1KdPaxK"
	contract := "TSNgPpTH2Pp7bYLct5m7TUmWrvw3JotYES"
	// 数量
	amount := decimal.NewFromFloat(0.3)
	var amountdecimal = decimal.New(1, 6)
	amountac, _ := amount.Mul(amountdecimal).Float64()

	parameter := util.ProcessBalanceOfParameter(addr, util.GetMethodID("ownerWithdraw(address,uint256)"), true, int64(amountac))
	transferContract, err := tronCli.TransferContract(pri, contract, parameter, 500000)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(transferContract)
}

func TestAddress(t *testing.T) {
	tronCli := NewTronCli("35.180.51.163:50051")
	err := tronCli.Torn.Start()
	if err != nil {
		fmt.Println("ERR: ", err.Error())
		return
	}
	/*
		Tron 地址0x41
	*/
	// Address转账户 0x8cfb6c7c01e60e28d8c8e9362bf6b0ad435db819
	with := util.AddressDealWith("0x8cfb6c7c01e60e28d8c8e9362bf6b0ad435db819")
	fmt.Println("With: ", with)
	s := address.HexToAddress(with).String()
	fmt.Println("Account: ", s)
}
