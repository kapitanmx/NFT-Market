package controllers

import (
	"NFTMarket/models"

	"fmt"
	"context"
	"net/http"
	"time"
	"math/rand"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func CreateNFT() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var nft model.NFT
		defer cancel()
		if err := c.BindJSON(&nft); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if validationErr := validate.Struct(NFT); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer cancel()


	}
}

func GetNFT() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		nftId := c.Param("nftId")
		var nft model.NFT
		defer cancel()
		if err := c.BindJSON(&nft); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result, err := nftCollection.FindOne(ctx, bson.M{"nft_id": nft})
		if err != nil {
			msg := fmt.Sprintf("Error occured while finding NFT")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func GetNFTS() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var nft model.NFT
		defer cancel()
	}
}


func GenerateUniqueHash() string {
	
}