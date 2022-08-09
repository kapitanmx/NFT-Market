package routes

import (
	"NFTMarket/controllers"
	"github.com/gin-gonic/gin"
)

func TransactionRoutes(router *gin.Engine) {
	router.GET("/transactions/", controllers.GetTransactions())
	router.GET("/transactions/:_id", controllers.GetTransactionById())
	router.GET("transactions/:date_from/:date_to", controllers.GetTransactionsByTimeInterval())
	router.POST("/transactions/", controllers.GenerateTransaction())
	router.PATCH("/transactions/:_id", controllers.EditTransaction())
	router.DELETE("/transactions/:_id", controllers.DeleteTransaction())
}

