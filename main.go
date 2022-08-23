package main

import (
	"os"

	"github.com/gin-gonic/gin"

	routes "NFTMarket/routes"
	

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.AnnouncementRoutes(router)
	routes.CartRoutes(router)
	routes.NftRoutes(router)
	routes.OrderRoutes(router)
	routes.TransactionRoutes(router)
	routes.WalletRoutes(router)

	router.Run(":" + port)
}
