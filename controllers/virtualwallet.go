package controllers

import (
	"NFTMarket/models"

	"fmt"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	walletCollection *mongo.Collection = db.GetCollection(db.DB, "wallets")
	userCollection *mongo.Collection = db.GetCollection(db.DB, "users")
)

func GenerateVirtualWallet() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var wallet models.VirtualWallet
		var user models.User
		userId = c.Param("user_id")
		defer cancel()
		if err := c.BindJSON(&wallet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			msg := Sprintf("User not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		result, err := walletCollection.InsertOne(ctx, wallet)
		defer cancel()
		if err != nil {
			msg := Sprintf("error: cannot create wallet")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, wallet)
	}
}

func GetWallet() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var wallet models.VirtualWallet
		walletId = c.Param("wallet_id")
		defer cancel()
		if err := c.BindJSON(&wallet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer cancel()
		err := walletCollection.FindOne(ctx, bson.M{"wallet_id": walletId}).Decode(&wallet)
		if err != nil {
			msg := Sprintf("Error: cannot get wallet")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, wallet)
		defer cancel()
	}
}

func DeleteWallet() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var wallet models.VirtualWallet
		walletId := c.Param("wallet_id")
		defer cancel()
		result, err := walletCollection.DeleteOne(ctx, bson.M{"wallet_id": walletId})
		if err != nil {
			msg := Sprintf("Error: cannot find current wallet")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, result)
		defer cancel()
	}
}