package routes

import (
	"NFTMarket/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	router.GET("/orders/", controllers.GetAllOrders())
	router.GET("/orders/:_id", controllers.GetOrder())
	router.POST("/orders/", controllers.MakeOrder())
	router.PATCH("/orders/:_id", controller.EditOrder())
	router.DELETE("/orders/:_id", controllers.DeleteOrder())
}