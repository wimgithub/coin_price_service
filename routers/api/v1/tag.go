package v1

import (
	"coin_price_service/pkg/app"
	"coin_price_service/pkg/e"
	"coin_price_service/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SigModel struct {
	Address string `json:"address"`
	Value   string `json:"value"`
}

func GetSig(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form SigModel
	)
	err := app.BindArgs(c, &form)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	sigM, err := service.GetSig(form.Address, form.Value)
	if err != nil {
		appG.Response(http.StatusOK, e.GET_SIG_ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, sigM)
}

func GetPrice(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	appG.Response(http.StatusOK, e.SUCCESS, app.GetPrice())
}

func GetPriceV2(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	appG.Response(http.StatusOK, e.SUCCESS, app.GetPriceV2())
}
