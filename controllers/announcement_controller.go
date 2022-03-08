package controllers

import (
	"NFTMarket/configs"
	"NFTMarket/models"
	"NFTMarket/responses"
	"context"
	"net/http"
	"proj1/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var announcementCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateAnnouncement() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var announcement models.Announcement
		defer cancel()

		if err := c.BindJSON(&announcement); err != nil {
			c.JSON(http.StatusBadRequest, responses.AnnouncementResponse{Status: http.StatusBadRequest})
			return
		}

		if validationErr := validate.Struct(&announcement); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AnnouncementResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newAnnouncement := models.Announcement{
			ID:             announcement.ID,
			Date:           announcement.Date,
			ExpDate:        announcement.ExpDate,
			Title:          announcement.Title,
			Desc:           announcement.Desc,
			Imgs:           announcement.Imgs,
			Price:          announcement.Price,
			AdvertiserName: announcement.AdvertiserName,
			AdvertiserID:   announcement.AdvertiserID,
			Category:       announcement.Category,
			Price:          announcement.Price,
			Tags:           announcement.Tags,
		}

		result, err := announcementCollection.InsertOne(ctx, newAnnouncement)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AnnouncementResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}

		c.JSON(http.StatusCreated, responses.AnnouncementResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAnnouncement() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		announcementId := c.Param("announcementId")
		var announcement models.Announcement
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(announcementId)
		err := announcementCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&announcement)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AnnouncementResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.AnnouncementResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": announcement}})

	}
}

func EditAnnouncement() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		announcementId := c.Param("announcementId")
		var announcement models.Announcement
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(announcementId)

		if err := c.BindJSON(&announcement); err != nil {
			c.JSON(http.StatusBadRequest, responses.AnnouncementResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AnnouncementResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validation.Error()}})
			return
		}

		update := bson.M{
			"date":            annoannouncement.Date,
			"exp_date":        announcement.ExpDate,
			"title":           announcement.Title,
			"desc":            announcement.Desc,
			"imgs":            announcement.Imgs,
			"price":           announcement.Price,
			"advertiser_name": announcement.AdvertiserName,
			"advertiser_id":   announcement.AdvertiserID,
			"category":        announcement.Category,
			"tags":            announcement.Tags,
		}

		result, err := announcementCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AnnouncementResponse{Status: http.StatussInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var updatedAnnouncement models.Announcement
		if result.MatchedCount == 1 {
			err := announcementCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedAnnouncement)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.AnnouncementResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			c.JSON(http.StatusOK, responses.AnnouncementResponse{Status: htttp.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedAnnouncement}})
		}
	}
}

func DeleteAnnouncement() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		announcementId := c.Param("announcementId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(announcementtId)

		result, err := announcementCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AnnouncementResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.AnnouncementResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Announcement with specified ID not found"}},
			)
			return
		}
		c.JSON(http.StatusOK,
			responses.AnnouncementResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Announcement successsfully deleted."}},
		)
	}
}

func GetAllAnnouncements() {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var announcement []models.Announcement
		defer cancel()

		result, err := announcementCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AnnouncementResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer result.Close(ctx)
		for result.Next(ctx) {
			var singleAnnouncement models.Announcement
			if err = result.Decode(&singleAnnouncement); err != nil {
				c.JSON(http.StatusInternalServerError, responses.AnnouncementResponse{status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}
			announcements = append(announcements, singleAnnouncement)
		}

		c.JSON(http.StatusOK,
			responses.AnnouncementResponse{status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": announcements}}
		)

	}
}
