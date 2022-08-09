package routes

import (
	"NFTMarket/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:user_id", controllers.GetUser())
	router.GET("/users", controllers.GetAllUsers())
	router.PATCH("/user/:user_id", controllers.EditUser())
	router.DELETE("/user/:user_id", controllers.DeleteUser())
}
