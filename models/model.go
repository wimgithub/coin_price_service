package model

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type RechargeEvents struct {
	gorm.Model
	RechargeTransactionHash string          `gorm:"unique_index:t_x"`
	BlockNumber             int64           // 充值高度
	BlockTimestamp          int64           // 充值时间
	Contract                string          // 充值合约地址
	UserAddress             string          // 用户
	Value                   decimal.Decimal `gorm:"column:value" sql:"type:decimal(32,16);"` // 充值金额
	ChainType               string          // 链类型 Ethereum || Tron
	FreedValue              decimal.Decimal `gorm:"default:0" sql:"type:decimal(32,16);"` // 已释放数量
}

type FreedRecord struct {
	gorm.Model
	Address    string          // 释放给谁
	FreedValue decimal.Decimal `sql:"type:decimal(32,16);"` // 释放数量
	Hash       string
	Status     int    // 1放币中/2放币成功/3放币失败
	ChainType  string // Ether || Tron
}

type Data struct {
	BlockNumber     int64  `json:"block_number"`
	BlockTimestamp  int64  `json:"block_timestamp"`
	ContractAddress string `json:"contract_address"`
	Result          Result `json:"result"`
	TransactionId   string `json:"transaction_id"`
	// ===account event
	From  string `json:"from"`
	To    string `json:"to"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Result struct {
	Addr  string `json:"addr"`
	Value string `json:"value"`
}

// 地址绑定
type BidAddress struct {
	gorm.Model
	TronAddress  string
	EtherAddress string
}

// Price
type HuoBiPrice struct {
	Status string `json:"status"`
	Tick   Ticker `json:"tick"`
}

type Ticker struct {
	Close decimal.Decimal `json:"close"`
}

type PriceResp struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}
