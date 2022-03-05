package api

import (
	"models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var user = models.User
var announcement = models.Announcement

var users = []user{}
var announcements = []announcement{}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.GET("/announcements", getAnnouncements)
	router.GET("/announcements/:id", getAnnouncementByID)
	router.POST("/users", postUsers)
	router.POST("/announcements", postAnnouncements)

	router.Run("localhost:8080")
}

func getUsers(c *gin.Context) {
	c.IndendedJSON(http.StatusOK, users)
}

func getAnnouncements(c *gin.Context) {
	c.IndendedJSON(http.StatusOK, announcement)
}

func postUsers(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func postAnnouncements(c *gin.Context) {
	var newAnnouncement announcement

	if err := c.BindJSON(&newAnnouncement); err != nil {
		return
	}

	announcements = append(announcements, newAnnouncement)
	c.IndentedJSON(http.StatusCreated, announcement)
}

func getUserByID(c *gin.Context) {
	id := c.Params("id")

	for _, a := range users {
		if a.ID == id {
			c.IntendedJSON(http.StatusOK, a)
			return
		}
	}
	c.IntendedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func getAnnouncementByID(c *gin.Context) {
	id := c.Params("id")

	for _, a := range announcements {
		if a.ID == id {
			c.IntendedJSON(http.StatusOK, a)
			return
		}
	}
	c.IntendedJSON(http.StatusNotFound, gin.H{"message": "announcement not found"})
}
