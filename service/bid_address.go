package service

import (
	"coin_price_service/pkg/conversion"
	"coin_price_service/pkg/rand"
	"coin_price_service/pkg/setting"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type SigResponseModel struct {
	To    string `json:"to"`
	Value string `json:"value"`
	Hash  string `json:"hash"`
	Sig   string `json:"sig"`
	Rand  string `json:"rand"`
}

func GetSig(addr, value string) (sigModel *SigResponseModel, err error) {
	privateKey, err := crypto.HexToECDSA(setting.EtherscanSetting.Private)
	if err != nil {
		return nil, err
	}

	var data []byte
	// to
	to := common.HexToAddress(addr)
	fmt.Println("TO: ", to.String())
	data = append(data, to.Bytes()...)

	// value
	wei := conversion.New().ToWei(value, 18)
	fmt.Println("WEI: ", wei)
	data = append(data, common.LeftPadBytes(wei.Bytes(), 32)...)

	// rand
	r := rand.GetRandomString(32)
	fmt.Println("RAND: ", r)
	data = append(data, []byte(r)...)

	hash := crypto.Keccak256Hash(data)
	fmt.Println("HASH: ", hash.Hex())

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}
	fmt.Println("SIG: ", hexutil.Encode(signature))

	return &SigResponseModel{
		To:    addr,
		Value: value,
		Hash:  hash.Hex(),
		Sig:   hexutil.Encode(signature),
		Rand:  r,
	}, nil

}
