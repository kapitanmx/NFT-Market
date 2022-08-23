package controllers

import (
	"NFTMarket/db"
	"context"
	"net/http"
	"NFTMarket/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var announcementCollection *mongo.Collection = db.GetCollection(db.DB, "users")
var validate = validator.New()

func CreateAnnouncement() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var announcement models.Announcement
		defer cancel()

		if err := c.BindJSON(&announcement); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(&announcement); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
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
			msg := Sprintf("Cannot create an announcement")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
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
			msg := Sprintf("Cannot get an announcement, maybe it caused by internal error or it doesn't even exist")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, result)

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
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		update := bson.M{
			"date":            announcement.Date,
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var updatedAnnouncement models.Announcement
		if result.MatchedCount == 1 {
			err := announcementCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedAnnouncement)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, updatedAnnouncement)
		}
	}
}

func DeleteAnnouncement() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		announcementId := c.Param("announcementId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(announcementId)

		result, err := announcementCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if result.DeletedCount < 1 {
			msg := Sprintf("Cannot found an announcement")
			c.JSON(http.StatusNotFound, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func GetAllAnnouncements() func(*gin.Context){
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var announcements []models.Announcement
		defer cancel()

		result, err := announcementCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		defer result.Close(ctx)

		for result.Next(ctx) {
			var singleAnnouncement models.Announcement
			if err = result.Decode(&singleAnnouncement); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			announcements = append(announcements, singleAnnouncement)
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}
