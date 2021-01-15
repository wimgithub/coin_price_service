package tron

import (
	"context"
	"coin_price_service/pkg/crypto"
	"coin_price_service/pkg/setting"
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/golang/protobuf/proto"
	"time"
)

type Tron interface {
	Start()
	TransferContract(ownerKey *ecdsa.PrivateKey, Contract string, data []byte, feeLimit int64) (string, error)
	GetConstantResultOfContract(from, Contract string, data []byte) ([][]byte, error)
}

type Cli struct {
	Torn *client.GrpcClient
}

func NewTronCli() Tron {
	return &Cli{Torn: client.NewGrpcClient(setting.TronscanSetting.Node)}
}

func (g *Cli) Start() {
	g.Start()
}

func timeoutContext() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	go func() {
		time.Sleep(time.Second * 60)
		cancel()
	}()
	return ctx
}

// Send
func (g *Cli) TransferContract(ownerKey *ecdsa.PrivateKey, Contract string, data []byte, feeLimit int64) (string, error) {
	transferContract := new(core.TriggerSmartContract)
	transferContract.OwnerAddress = crypto.PubkeyToAddress(ownerKey.
		PublicKey).Bytes()
	transferContract.ContractAddress, _ = common.DecodeCheck(Contract)
	transferContract.Data = data
	transferTransactionEx, err := g.Torn.Client.TriggerConstantContract(timeoutContext(), transferContract)
	var txid string
	if err != nil {
		return txid, err
	}
	transferTransaction := transferTransactionEx.Transaction
	if transferTransaction == nil || len(transferTransaction.
		GetRawData().GetContract()) == 0 {
		return txid, fmt.Errorf("transfer error: invalid transaction")
	}
	if feeLimit > 0 {
		transferTransaction.RawData.FeeLimit = feeLimit
	}

	hash, err := g.signTransaction(transferTransaction, ownerKey)
	if err != nil {
		return txid, err
	}
	txid = hexutil.Encode(hash)

	result, err := g.Torn.Client.BroadcastTransaction(timeoutContext(),
		transferTransaction)
	if err != nil {
		return "", err
	}
	if !result.Result {
		return "", fmt.Errorf("api get false the msg: %v", result.String())
	}
	return txid, err
}

// Call
func (g *Cli) GetConstantResultOfContract(from, Contract string, data []byte) ([][]byte, error) {
	transferContract := new(core.TriggerSmartContract)
	transferContract.OwnerAddress, _ = common.DecodeCheck(from)
	transferContract.ContractAddress, _ = common.DecodeCheck(Contract)
	transferContract.Data = data
	transferTransactionEx, err := g.Torn.Client.TriggerConstantContract(timeoutContext(), transferContract)
	if err != nil {
		return nil, err
	}
	if transferTransactionEx == nil || len(transferTransactionEx.GetConstantResult()) == 0 {
		return nil, fmt.Errorf("GetConstantResult error: invalid TriggerConstantContract")
	}
	return transferTransactionEx.GetConstantResult(), err
}

// SignTransaction 签名交易
func (g *Cli) signTransaction(transaction *core.Transaction, key *ecdsa.PrivateKey) ([]byte, error) {
	transaction.GetRawData().Timestamp = time.Now().UnixNano() / 1000000
	rawData, err := proto.Marshal(transaction.GetRawData())
	if err != nil {
		return nil, err
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	contractList := transaction.GetRawData().GetContract()
	for range contractList {
		signature, err := crypto.Sign(hash, key)
		if err != nil {
			return nil, err
		}
		transaction.Signature = append(transaction.Signature, signature)
	}
	return hash, nil
}
