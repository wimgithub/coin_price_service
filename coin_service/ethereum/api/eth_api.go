package api

import (
	"context"
	model "coin_price_service/models"
	"coin_price_service/models/mysql"
	"coin_price_service/pkg/logging"
	"coin_price_service/pkg/setting"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/karalabe/go-ethereum/crypto/sha3"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Ethereum interface {
	GetEthCli() *ethclient.Client
	GetAccountBalance(address string) string
	CreatEthAccount() (string, string, string)
	TransferERC20PrivateKey(privateKeyStr string, gasPrice *big.Int, gasLimit uint64, value *big.Int, to, token string) (error, string, string)
	GetErc20Balance(contractAddress, address string) *big.Int
	StartTransactionReceipt(hash, key string, record *model.RechargeEvents, value decimal.Decimal)
	TransactionReceipt(hash string) chan bool
	IsBalance(addr common.Address, gasPrice *big.Int, gasLimit *big.Int) bool
}

type ethereum struct {
	EthCli   *ethclient.Client
	nonceMap sync.Map
}

var LastNonce uint64
var mutex = &sync.Mutex{}

func NewEtherService() (Ethereum, error) {
	cli, err := ethclient.Dial(setting.EtherscanSetting.Infura)
	return &ethereum{
		EthCli: cli,
	}, err
}

func (e *ethereum) GetEthCli() *ethclient.Client {
	return e.EthCli
}

// 查询eth余额
func (e *ethereum) GetAccountBalance(address string) string {

	// 创建ETH客户端
	client := e.EthCli

	// 查询账户余额
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		logging.Error("账户查询失败:%s", err)
		return ""
	}
	logging.Info("", "WEI: ", balance)

	// Wei 转换为 ETH
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	logging.Info("", "ETH: ", ethValue)

	return ethValue.String()
}

// 创建以太坊公钥私钥地址
func (e *ethereum) CreatEthAccount() (string, string, string) {
	// 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		logging.Error("%s", err)
		return "", "", ""
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyStr := hexutil.Encode(privateKeyBytes)[2:]
	logging.Info("", "-String--private key ---", privateKeyStr)
	// 生成公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logging.Error("error casting public key to ECDSA")
		return "", "", ""
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyStr := hexutil.Encode(publicKeyBytes)[4:]
	logging.Info("", "-String--public key ---", publicKeyStr)
	// 生成地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	logging.Info("", "-String--address ---", address)

	return privateKeyStr, publicKeyStr, address
}

//ERC20 转账（privateKey）
func (e *ethereum) TransferERC20PrivateKey(privateKeyStr string, gasPrice *big.Int, gasLimit uint64, value *big.Int, to, token string) (error, string, string) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		logging.Error("私钥转换失败")
		return err, "", ""
	}
	return e.transactionErc20(privateKey, gasPrice, gasLimit, value, to, token)
}

func (e *ethereum) transactionErc20(privateKey *ecdsa.PrivateKey, gasPrice *big.Int, gasLimit uint64, amount *big.Int, to, token string) (error, string, string) {
	mutex.Lock()
	client := e.EthCli

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		mutex.Unlock()
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey"), "", ""
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	b := e.IsBalance(fromAddress, gasPrice, big.NewInt(int64(gasLimit)))
	if !b {
		return errors.New("手续费不足!"), "", ""
	}
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		logging.Error("%s", err)
		mutex.Unlock()
		return errors.New(err.Error()), "", ""
	}

	fromNonce := e.loadMap(fromAddress.String())
	logging.Info("当前提币账户本地维护的nonce值为 ：", fromNonce, "   最新nonce值为：", nonce)
	if fromNonce >= nonce+1 && nonce != 0 {
		mutex.Unlock()
		return errors.New("网络拥堵!"), "", ""
	}

	if fromNonce == 0 || fromNonce < nonce {
		e.storeMap(fromAddress.String(), nonce)
	}

	toAddress := common.HexToAddress(to)
	tokenAddress := common.HexToAddress(token)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewKeccak256()
	_, err = hash.Write(transferFnSignature)
	if err != nil {
		mutex.Unlock()
		return errors.New(err.Error()), "", ""
	}

	methodID := hash.Sum(nil)[:4]
	logging.Info("", "methodID:"+hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	logging.Info("", "paddedAddress:"+hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	logging.Info("", "paddedAmount:"+hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	logging.Info("GasPrice: ", gasPrice)
	tx := types.NewTransaction(nonce, tokenAddress, big.NewInt(0), gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		logging.Error("%s", err)
		mutex.Unlock()
		return errors.New(err.Error()), "", ""
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		logging.Error("%s", err)
		mutex.Unlock()
		return errors.New(err.Error()), "", ""
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		logging.Error("%s", err)
		return errors.New(err.Error()), "", ""
	}

	// 更新当前账户本地nonce
	e.storeMap(fromAddress.String(), e.loadMap(fromAddress.String())+1)
	logging.Info("ERC20 转账成功", "tx sent: %s", signedTx.Hash().Hex())

	mutex.Unlock()
	return nil, signedTx.Hash().Hex(), fromAddress.String()
}

func (e *ethereum) loadMap(key string) uint64 {
	load, ok := e.nonceMap.Load(key)
	if !ok {
		return 0
	}
	return load.(uint64)
}

func (e *ethereum) storeMap(key string, value uint64) {
	e.nonceMap.Store(key, value)
}

// 获取某账户的erc20代币余额， 如果 etherscan api挂掉可以使用此方法
func (e *ethereum) GetErc20Balance(contractAddress, address string) *big.Int {

	url := setting.EtherscanSetting.Infura

	payload := strings.NewReader("{\"jsonrpc\": \"2.0\",\"method\": \"eth_call\",\"params\": [{\"from\": \"" + address + "\",\"to\": \"" + contractAddress + "\",\"data\": \"0x70a08231000000000000000000000000" + address[2:] + "\"}, \"latest\"],\"id\":1}")

	i := 0
	var body []byte
	for {
		req, err := http.NewRequest("POST", url, payload)
		res, err := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, err = ioutil.ReadAll(res.Body)

		if i > 5 || err == nil {
			if err != nil {
				logging.Error("重试5次获取账户Erc20代币余额失败:%s", err)
			}
			break
		}
		time.Sleep(1 * time.Second)
		i = i + 1
	}
	balance := make(map[string]interface{})
	err := json.Unmarshal(body, &balance)
	if err != nil {
		logging.Error("获取erc20代币余额结构解析失败:%s", err)
	}
	if balance["error"] != nil {
		return big.NewInt(0)
	}
	n := new(big.Int)
	n, _ = n.SetString(balance["result"].(string)[2:], 16)
	return n
}

// 监听交易
func (e *ethereum) TransactionReceipt(hash string) chan bool {
	client := e.EthCli
	b := make(chan bool)
	go func() {
		for {
			receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(hash))
			if err != nil {
				if err.Error() == "not found" {
					time.Sleep(60 * time.Second)
					continue
				}
				logging.Error("Get TransactionReceipt err:%s", err)
				b <- false
				break
			}
			if receipt == nil {
				time.Sleep(60 * time.Second)
				continue
			}
			if receipt.Status == 1 {
				b <- true
				break
			} else if receipt.Status == 0 {
				b <- false
				break
			}
		}
	}()
	return b
}

func (e *ethereum) StartTransactionReceipt(hash, key string, record *model.RechargeEvents, value decimal.Decimal) {
	b := e.TransactionReceipt(hash)
	select {
	case t := <-b:
		i := 0
		if t == true {
			// 成功
			i = 2
			record.FreedValue = record.FreedValue.Add(value)
			_ = mysql.SharedStore().UpdateTxStatus(record.RechargeTransactionHash, record)
		} else {
			// 失败
			e.storeMap(key, e.loadMap(key)-1)
			i = 3
		}
		// 更新状态
		_ = mysql.SharedStore().UpdateFreedTxStatus(hash, &model.FreedRecord{Hash: hash, Status: i})
	}
}

// 检查余额
func (e *ethereum) IsBalance(addr common.Address, gasPrice *big.Int, gasLimit *big.Int) bool {
	// 检查余额
	balance, _ := e.EthCli.BalanceAt(context.Background(), addr, nil)
	gas := big.NewInt(0)
	gas.Mul(gasPrice, gasLimit)
	if balance.Cmp(gas) == 1 {
		return true
	}
	return false
}
