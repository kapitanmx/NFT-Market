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

func GenerateTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var transaction models.Transaction
		defer cancel()
	}
}

func GetTransactionById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		transactionId := c.Param("transactionId")
		var transaction models.Transaction
		defer cancel()
	}
}

func GetTransactionsByTimeInterval() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var transaction models.Transaction
		defer cancel()
	}
}

func GetTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var transaction models.Transaction
		defer cancel()
	}
}

func EditTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		transactionId := c.Param("transactionId")
		var transaction models.Transaction
		defer cancel()
	}
}

func DeleteTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		transactionId := c.Param("transactionId")
		var transaction models.Transaction
		defer cancel()
	}
}