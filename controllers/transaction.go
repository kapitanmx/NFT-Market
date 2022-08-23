package controllers

import (
	"NFTMarket/models"

	"fmt"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userCollection *mongo.Collection = db.GetCollection(db.DB, "users")
	transactionCollection *mongo.Collection = db.GetCollection(db.DB, "transactions")
)

func GenerateTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var transaction models.Transaction
		var user models.User
		userId := c.Param("user_id")
		defer cancel()
		if err := c.BindJSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if validationErr := validate.Struct(transaction); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := userCollection.FindOne(ctx, bson.M{"user_id" : userId})
		defer cancel()
		if err != nil {
			msg := Sprintf("User not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		result, err := transactionCollection.InsertOne(ctx, transaction)
		defer cancel()
		if err != nil {
			msg := Sprintf("Unable to add new transaction")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, transaction)
	}
}

func GetTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		transactionId := c.Param("transactionId")
		var transaction models.Transaction
		defer cancel()
		
	}
}

func GetTransactionsByTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var transaction models.Transaction
		defer cancel()
	}
}

func GetTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		results, err := transactionList.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			msg := Sprintf("An error occured while listing transactions")
			c.JSON(htttp.StatusBadRequest, gin.H{"error" : msg})
			return
		}
		var transactions []bson.M
		if err = results.All(ctx, &transactions); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, transactions)
	}
}

func EditTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		transactionId := c.Param("transactionId")
		var transaction models.Transaction
		defer cancel()
		if err := c.BindJSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		defer cancel()
		updateObj := primitive.D
		result, err := transactionCollection.UpdateOne(ctx, bson.M{"transaction_id": transactionId}, bson.M{"$set": updateObj})
		if err != nil {
			msg := Sprintf("Unable to edit transaction")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, transaction)
		defer cancel()
	}
}

func DeleteTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		transactionId := c.Param("transactionId")
		var transaction models.Transaction
		defer cancel()
		result, err := transactionCollection.DeleteOne(ctx, bson.M{"transaction_id": transactionId})
		if err != nil {
			msg := Sprintf("Error: cannot delete transaction")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, transaction)
		defer cancel()
	}
}