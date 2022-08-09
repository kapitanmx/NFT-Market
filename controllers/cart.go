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

	}
}

func ModifyCart() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func DeleteCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}