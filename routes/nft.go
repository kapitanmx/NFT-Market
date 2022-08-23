package routes

import (
	"NFTMarket/controllers"
	"github.com/gin-gonic/gin"
)

func NftRoutes(router *gin.Engine) {
	router.GET("/nfts/", controller.GetNFTS())
	router.GET("/nfts/:_id", controller.GetNFT())
	router.POST("/nfts/", controller.CreateNFT())
}