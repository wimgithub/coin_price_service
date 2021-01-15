package routers

import (
	_ "coin_price_service/docs"
	"coin_price_service/middleware/core"
	"coin_price_service/routers/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(core.Cors())
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/get_address", v1.GetBidAddress)
		apiv1.POST("/bid_address", v1.Bid)
	}

	return r
}
