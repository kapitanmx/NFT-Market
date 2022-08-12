package routes

import (
	"NFTMarket/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/users/", controllers.SignUp())
	router.GET("/users/:user_id", controllers.GetUser())
	router.GET("/users/", controllers.GetAllUsers())
	router.PATCH("/users/:user_id", controllers.EditUser())
	router.DELETE("/users/:user_id", controllers.DeleteUser())
}
