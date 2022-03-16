package routes

import (
	"NFTMarket/controllers"

	"github.com/gin-gonic/gin"
)

func AnnouncementRouter(router *gin.Engine) {
	router.POST("/announcement", controllers.CreateAnnouncement())
	router.GET("/announcement/:announcementId", controllers.GetAnnouncement())
	router.GET("/announcements", controllers.GetAllAnnouncements())
	router.PUT("/announcement/:announcementId", controllers.EditAnnouncement())
	router.DELETE("/announcement/:announcementId", controllers.DeleteAnnouncement())
}
