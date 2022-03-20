package models

import (
	"crypto/sha256"
	"io"
	"os"

	"github.com/google/uuid"
)

type NFT struct {
	TokenID          string
	CreationDate     string
	TokenName        string
	TokenHashCode    []byte
	DigitalSignature string
	Img              string
}

func CreateNFT(
	id,
	creationDate,
	name,
	digitalSignature,
	img string,
) (*NFT, error) {
	id, err := uuid.UUID()
	if err != nil {
		return nil, err
	}
	hashCode, err := generateHash()
	return &NFT{
		TokenID:          id,
		CreationDate:     creationDate,
		TokenName:        name,
		TokenHashCode:    hashCode,
		DigitalSignature: digitalSignature,
		Img:              img,
	}, nil
}

func generateHash() ([]byte, error) {
	env := os.Getenv("KEYS")
	file, err := os.Open(env)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := make([]byte, 30*1024)
	sha256 := sha256.New()
	for {
		n, err := file.Read(buf)
		if n > 0 {
			_, err := sha256.Write(buf[:n])
			if err != nil {
				return nil, err
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}
	}
	sum := sha256.Sum(nil)
	return sum, nil
}

func setDigitalSignature() {}
