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
		"bsvusdt": "BSV",
		"htusdt":  "HT",
		"filusdt": "FIL",
		"ethusdt": "ETH",
		"btcusdt": "BTC",
		"ltcusdt": "LTC",
		"bchusdt": "BCH",
		"dotusdt": "DOT",
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
