package service

import (
	model "coin_price_service/models"
	"coin_price_service/models/mysql"
	"coin_price_service/mysql_service"
	"coin_price_service/pkg/e"
)

func BidAddress(tron, ether string) int {
	b := mysql_service.IsBidAddress(tron, ether)
	if b == false {
		return e.BID_REPEAT
	}
	err := mysql.SharedStore().BidInsert(&model.BidAddress{TronAddress: tron, EtherAddress: ether})
	if err != nil {
		return e.BID_REPEAT
	}
	return e.SUCCESS
}

func GetBidAddress(addr, t string) string {
	var bid []*model.BidAddress
	if t == "1" {
		bid, _ = mysql.SharedStore().QueryBid(&model.BidAddress{TronAddress: addr})
		if len(bid) > 0 {
			return bid[0].EtherAddress
		}
	} else if t == "2" {
		bid, _ = mysql.SharedStore().QueryBid(&model.BidAddress{EtherAddress: addr})
		if len(bid) > 0 {
			return bid[0].TronAddress
		}
	}
	return ""
}
