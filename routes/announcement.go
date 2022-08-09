package routes

import (
	"NFTMarket/controllers"
	"github.com/gin-gonic/gin"
)

func AnnouncementRouter(router *gin.Engine) {
	router.POST("/announcement", controllers.CreateAnnouncement())
	router.GET("/announcement/:_id", controllers.GetAnnouncement())
	router.GET("/announcements", controllers.GetAllAnnouncements())
	router.PUT("/announcement/:_id", controllers.EditAnnouncement())
	router.DELETE("/announcement/:_id", controllers.DeleteAnnouncement())
}
