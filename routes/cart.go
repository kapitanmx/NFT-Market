package routes

import (
	"NFTMarket/controllers"
	"github.com/gin-gonic/gin"
)

func CartRoutes(router *gin.Engine) {
	router.POST("/cart/", controllers.CreateCart())
	router.GET("/cart/:cart_id", controllers.GetCart())
	router.PATCH("/cart/:cart_id", controllers.ModifyCart())
	router.DELETE("/cart/:cart_id", controllers.DeleteCart())
}