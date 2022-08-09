package controllers

import (
	"fmt"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validator = validator.New()

func CreateCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var cart models.Cart
		defer cancel()
	}
}

func ModifyCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cartId := c.Param("cartId")
		var cart models.Cart
		defer cancel()
	}
}

func DeleteCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cartId := c.Param("cartId")
		var cart models.Cart
		defer cancel()
	}
}