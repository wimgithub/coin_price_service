package v1

import (
	"coin_price_service/pkg/app"
	"coin_price_service/pkg/e"
	"coin_price_service/pkg/util"
	"coin_price_service/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddTagForm struct {
	Tron    string `json:"tron"`
	Ether   string `json:"ether"`
	Address string `json:"address"`
	Type    string `json:"type"` // 1 tron, 2 ether
}

func Bid(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddTagForm
	)
	err := app.BindArgs(c, &form)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	// Tron / Ether 地址验证
	if util.TronValidAddress(form.Tron) == false {
		appG.Response(http.StatusOK, e.TRON_ADDR_ERROR, nil)
		return
	}
	if util.EtherValidAddress(form.Ether) == false {
		appG.Response(http.StatusOK, e.ETHER_ADDR_ERROR, nil)
		return
	}
	status := service.BidAddress(form.Tron, form.Ether)
	if status != e.SUCCESS {
		appG.Response(http.StatusBadRequest, status, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetBidAddress(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddTagForm
	)
	err := app.BindArgs(c, &form)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, service.GetBidAddress(form.Address, form.Type))
}
