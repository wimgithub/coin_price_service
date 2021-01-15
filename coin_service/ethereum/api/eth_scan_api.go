package api

import (
	"coin_price_service/pkg/setting"
	"github.com/wimgithub/etherscan-api"
)

type EthereumScan interface {
	GetEthScanCli() *etherscan.Client
	GetErc20Balance(contractAddress, address string) (*etherscan.BigInt, error)
}

type ethereumScan struct {
	EthScanCli *etherscan.Client
}

func NewEthereumScanService() EthereumScan {
	var netWork etherscan.Network
	// 主网: Mainnet 测试网: Ropsten Kovan Rinkby Tobalaba
	switch setting.EtherscanSetting.NetWork {
	case "Mainnet":
		netWork = etherscan.Mainnet
		break
	case "Ropsten":
		netWork = etherscan.Ropsten
		break
	case "Kovan":
		netWork = etherscan.Kovan
		break
	case "Rinkby":
		netWork = etherscan.Rinkby
		break
	case "Tobalaba":
		netWork = etherscan.Tobalaba
		break
	}

	// 创建连接指定网络的客户端
	scanCli := etherscan.New(netWork, setting.EtherscanSetting.EthScanAPIKey)
	return &ethereumScan{
		EthScanCli: scanCli,
	}
}

func (e *ethereumScan) GetEthScanCli() *etherscan.Client {
	return e.EthScanCli
}

func (e *ethereumScan) GetErc20Balance(contractAddress, address string) (*etherscan.BigInt, error) {
	balance, err := e.EthScanCli.TokenBalance(contractAddress, address)
	return balance, err
}
