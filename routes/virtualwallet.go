package routes

import (
	"NFTMarket/controllers"
	"github.com/gin-gonic/gin"
)

func WalletRoutes(router *gin.Engine) {
	router.POST("/virtualwallet/", controllers.GenerateVirtualWallet())
	router.GET("/virtualwallet/:_id", controllers.GetWallet())
	router.PATCH("/virtualwallet/:_id", controllers.ModifyWallet())
}
