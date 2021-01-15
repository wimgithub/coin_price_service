package main

import (
	model "coin_price_service/models"
	"coin_price_service/pkg/http_util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"coin_price_service/pkg/gredis"
	"coin_price_service/pkg/logging"
	"coin_price_service/pkg/setting"
	"coin_price_service/pkg/util"
	"coin_price_service/routers"
)

func init() {
	setting.Setup()
	logging.Setup()
	util.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://coin_price_service
// @license.name MIT
// @license.url https://coin_price_service/blob/master/LICENSE
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	GetPrice2()
	_ = server.ListenAndServe()
	signalNotifyExit()
}

func signalNotifyExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		logging.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			close(gredis.RedisSnapshot)
			logging.Info("coin_price_service exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func GetPrice(c *gin.Context) {
	var coins = []string{"bsvusdt", "htusdt", "filusdt", "ethusdt", "btcusdt", "ltcusdt", "bchusdt", "dotusdt"}
	url := "https://api.huobi.pro/market/detail/merged?symbol="
	var PData *model.HuoBiPrice
	for _, v := range coins {
		fmt.Println("price: ", url+v)
		bytes, _ := http_util.Get(url + v)
		_ = json.Unmarshal(bytes, &PData)
		fmt.Println(v, ": ", PData.Tick.Close)
	}
}

func GetPrice2() {
	var coins = []string{"bsvusdt", "htusdt", "filusdt", "ethusdt", "btcusdt", "ltcusdt", "bchusdt", "dotusdt"}
	url := "https://api.huobi.pro/market/detail/merged?symbol="
	var PData model.HuoBiPrice
	ch := make(chan []byte, len(coins))
	for _, v := range coins {
		go func(n string) {
			fmt.Println("name: ", n)
			bytes, err := http_util.Get(url + n)
			if err != nil {
				fmt.Println(err)
				return
			}
			ch <- bytes[:]
		}(v)
	}
	for {
		select {
		case data := <-ch:
			if len(data) == 0 {
				break
			}
			json.Unmarshal(data, &PData)
			fmt.Println("Close: ", PData.Tick.Close)
		}
		break
	}
}
