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
	productCollection *mongo.GetCollection(db.DB, "products")
	cartCollection *mongo.GetCollection(db.DB, "carts")
)

func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var product models.Product
		defer cancel()

		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if validationErr := validate.Struct(product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := cartCollection.FindOne(ctx, bson.M{"cart_id": cart.ID}).Decode(&cart)
		defer cancel()
		if err != nil {
			msg := fmt.Sprintf("Product not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return 
		}

		product.ID := primitive.NewObjectID()
		result, err := productCollection.InsertOne(ctx, product)
		if err != nil {
			msg := fmt.Sprintf("Product not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, product)
	}
}

func GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		productId := c.Param("_id")
		var product models.Product
		defer cancel()

		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := productCollection.FindOne(ctx, bson.M{"product_id": productId})
		if err != nil {
			msg := Sprintf("Product not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

func GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		productId := c.Param("_id")
		var product models.Product
		defer cancel()
		
	}
}

func DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		productId := c.Param("_id")
		var product models.Product
		defer cancel()
	}
}