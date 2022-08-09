package models

import (
	"crypto/sha256"
	"io"
	"os"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NFT struct {
	TokenID          primitive.ObjectID `bson:"_id"` 
	CreationDate     string `json:"creation_date" validation:"required"`
	TokenName        string `json:"token_name" validation:"required"`
	TokenHashCode    []byte `json:"token_hash_code" validation:"required"`
	DigitalSignature string `json:"digital_signature" validation:"required"`
	Img              string `json:"img" validation:"required"`
}
