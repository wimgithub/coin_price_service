package mysql_service

import (
	model "coin_price_service/models"
	"coin_price_service/models/mysql"
)

// 验证是否绑定
func IsBidAddress(tron, ether string) bool {
	tronBid, _ := mysql.SharedStore().QueryBid(&model.BidAddress{TronAddress: tron})
	if len(tronBid) != 0 {
		return false
	}
	etherBid, _ := mysql.SharedStore().QueryBid(&model.BidAddress{EtherAddress: ether})
	if len(etherBid) != 0 {
		return false
	}
	return true
}
