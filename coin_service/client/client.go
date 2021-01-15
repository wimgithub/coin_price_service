package client

import (
	eth_api "coin_price_service/coin_service/ethereum/api"
	tron_api "coin_price_service/coin_service/tron"
	"coin_price_service/pkg/logging"
	"fmt"
)

var (
	EthCli     eth_api.Ethereum
	EthScanCli eth_api.EthereumScan
	TronCli    tron_api.Tron
)

func NewCliService() {
	NewEthServer()
	NewTronServer()
}

func NewEthServer() {
	var err error
	EthCli, err = eth_api.NewEtherService()
	if err != nil {
		logging.Error("eth client creat err: ", err)
	}
	fmt.Println("EtherEumCli创建成功...")

	EthScanCli = eth_api.NewEthereumScanService()
	fmt.Println("EtherScanCli创建成功...")
}

func NewTronServer() {
	TronCli = tron_api.NewTronCli()
}
