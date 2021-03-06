package app

import (
	model "coin_price_service/models"
	"coin_price_service/pkg/http_util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func BindArgs(c *gin.Context, obj interface{}) error {
	if err := c.Bind(&obj); err != nil {
		c.Abort()
		return err
	}
	if err := binding.Validator.ValidateStruct(obj); err != nil {
		c.Abort()
		return err
	}
	return nil
}

/*
返回值取 close
    BSV-USDT
    API：https://api.huobi.pro/market/detail/merged?symbol=bsvusdt
    HT-USDT
    API：https://api.huobi.pro/market/detail/merged?symbol=htusdt
    FIL-USDT
    API：https://api.huobi.pro/market/detail/merged?symbol=filusdt
    ETH-USDT
    API：https://api.huobi.pro/market/detail/merged?symbol=ethusdt
    BTC-USDT
    API：https://api.huobi.pro/market/detail/merged?symbol=btcusdt
    LTC-USDT
    API：https://api.huobi.pro/market/detail/merged?symbol=ltcusdt
    BCH-USDT
    API：https://api.huobi.pro/market/detail/merged?symbol=bchusdt
    DOT-USDT
    API：https://api.huobi.pro/market/detail/merged?symbol=dotusdt
*/
func GetPrice() (price []*model.PriceResp) {
	var coins = map[string]string{
		"bsvusdt": "HBSVHTCPool",
		"htusdt":  "HTHTCPool",
		"filusdt": "HFILHTCPool",
		"ethusdt": "HETHHTCPool",
		"btcusdt": "HBTCHTCPool",
		"ltcusdt": "HLTCHTCPool",
		"bchusdt": "HBCHHTCPool",
		"dotusdt": "HDOTHTCPool",
	}
	url := "https://api.huobi.pro/market/detail/merged?symbol="
	var PData *model.HuoBiPrice

	for k, v := range coins {
		fmt.Println("price: ", url+k)
		bytes, _ := http_util.Get(url + k)
		_ = json.Unmarshal(bytes, &PData)
		price = append(price, &model.PriceResp{
			Name:  v,
			Price: PData.Tick.Close.String(),
		})
	}
	price = append(price, &model.PriceResp{
		Name:  "HUSDHTCPool",
		Price: "1.0",
	})
	return
}

func GetPriceV2() (price []*model.PriceResp) {
	var coins = map[string]string{
		"bsvusdt": "HBSV",
		"htusdt":  "HT",
		"filusdt": "HFIL",
		"ethusdt": "HETH",
		"btcusdt": "HBTC",
		"ltcusdt": "HLTC",
		"bchusdt": "HBCH",
		"dotusdt": "HDOT",
	}
	url := "https://api.huobi.pro/market/detail/merged?symbol="
	var PData *model.HuoBiPrice

	for k, v := range coins {
		fmt.Println("price: ", url+k)
		bytes, _ := http_util.Get(url + k)
		_ = json.Unmarshal(bytes, &PData)
		price = append(price, &model.PriceResp{
			Name:  v,
			Price: PData.Tick.Close.String(),
		})
	}
	price = append(price, &model.PriceResp{
		Name:  "HUSD",
		Price: "1.0",
	})
	price = append(price, &model.PriceResp{
		Name:  "USDT",
		Price: "1.0",
	})
	return
}
