package routes

import (
	"NFTMarket/controllers"
	"github.com/gin-gonic/gin"
)

func TransactionRoutes(router *gin.Engine) {
	router.GET("/transactions/", controllers.GetTransactions())
	router.GET("/transactions/:transaction_id", controllers.GetTransaction())
	router.GET("transactions/:date_from/:date_to", controllers.GetTransactionsByTime())
	router.POST("/transactions/", controllers.GenerateTransaction())
	router.PATCH("/transactions/:transaction_id", controllers.EditTransaction())
	router.DELETE("/transactions/:transaction_id", controllers.DeleteTransaction())
}

