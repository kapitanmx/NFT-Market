package routes

import (
	"NFTMarket/controllers"
	"github.com/gin-gonic/gin"
)

func AnnouncementRouter(router *gin.Engine) {
	router.POST("/announcements/", controllers.CreateAnnouncement())
	router.GET("/announcements/:_id", controllers.GetAnnouncement())
	router.GET("/announcements/", controllers.GetAllAnnouncements())
	router.PUT("/announcements/:_id", controllers.EditAnnouncement())
	router.DELETE("/announcements/:_id", controllers.DeleteAnnouncement())
}
