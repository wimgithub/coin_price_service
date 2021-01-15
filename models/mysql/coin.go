package mysql

import (
	"coin_price_service/models"
	"fmt"
	"strings"
)

func (s *Store) AddRecTxs(txs []*model.RechargeEvents) error {
	if len(txs) == 0 {
		return nil
	}
	var valueStrings []string
	for _, tx := range txs {
		valueString := fmt.Sprintf("(NOW(),'%v', %v, %v,'%v','%v',%v,'%v')",
			tx.RechargeTransactionHash, tx.BlockNumber, tx.BlockTimestamp, tx.Contract, tx.UserAddress, tx.Value, tx.ChainType)
		valueStrings = append(valueStrings, valueString)
	}
	sql := fmt.Sprintf("INSERT IGNORE INTO chain_recharge_events (created_at,recharge_transaction_hash,block_number,block_timestamp,contract,user_address,value,chain_type) VALUES %s",
		strings.Join(valueStrings, ","))
	return s.db.Exec(sql).Error
}

// 根据hash更新充值交易
func (s *Store) UpdateTxStatus(hash string, data *model.RechargeEvents) error {
	return s.db.Model(&model.RechargeEvents{}).Where("recharge_transaction_hash = ?", hash).Updates(&data).Error
}

// 根据hash更新释放交易
func (s *Store) UpdateFreedTxStatus(hash string, data *model.FreedRecord) error {
	return s.db.Model(&model.FreedRecord{}).Where("hash = ?", hash).Updates(&data).Error
}

// 获取充值信息
func (s *Store) GetFullRecord(t string) (txs []*model.RechargeEvents, err error) {
	err = s.db.Model(&model.RechargeEvents{}).Where("chain_type = ?", t).Having("value > freed_value").Find(&txs).Error
	return
}

func (s *Store) BidInsert(data *model.BidAddress) error {
	return s.db.Save(data).Error
}

func (s *Store) QueryBid(bid *model.BidAddress) (data []*model.BidAddress, err error) {
	err = s.db.Where(bid).First(&data).Error
	return
}

func (s *Store) FreedRecordInsert(data *model.FreedRecord) error {
	return s.db.Save(data).Error
}
