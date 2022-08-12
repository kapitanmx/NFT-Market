package controllers

import (
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
	cartCollection *mongo.Collection = db.GetCollection(db.DB, "carts")
	userCollection *mongo.Collection = db.GetCollection(db.DB, "users")
)
var validator = validator.New()

func CreateCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var cart models.Cart
		var user models.User
		defer cancel()

		if err := c.BindJSON(&cart); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error", err.Error()})
			return
		}

		if validationErr := validate.Struct(cart); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error", err.Error()})
			return
		}
		err := userCollection.FindOne(ctx, bson.M{"user_id": cart_id}).Decode(&user)
		defer cancel()
		if err != nil {
			msg := Sprintf("User not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		cart.ID = primitive.NewObjectID()
		var num = toFixed(*cart.TotalPrice, 2)
		cart.TotalPrice = &num

		result, err := cartCollection.InsertOne(ctx, cart)
		if err != nil {
			msg := Sprintf("Cart not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func ModifyCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cartId := c.Param("cartId")
		var cart models.Cart
		var user models.User
		var product models.Product
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(cartId)
		if err := c.BindJSON(&cart); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateObj primitive.D

		if cart.Products != nil {
			updateObj = append(updateObj, bson.E{"products", cart.Products})
		}

		if cart.TotalPrice != nil {
			updateObj = append(updateObj, bson.E{"total_price", cart.TotalPrice})
		}

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