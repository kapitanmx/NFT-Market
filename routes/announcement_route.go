package routes

import (
	"NFTMarket/controllers"

	"github.com/gin-gonic/gin"
)

func AnnouncementRouter(router *gin.Engine) {
	router.POST("/announcement", controllers.CreateAnnouncement())
}
