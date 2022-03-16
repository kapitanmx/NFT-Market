package models

import (
	"crypto/sha256"
	"io"
	"log"
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

type Hash uint

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
	nft := &NFT{
		TokenID:          id,
		CreationDate:     creationDate,
		TokenName:        name,
		TokenHashCode:    hashCode,
		DigitalSignature: digitalSignature,
		Img:              img,
	}
	return nft, nil
}

func generateHash() ([]byte, error) {
	file, err := os.Open("keys.txt")
	if err != nil {
		log.Fatal(err)
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
				log.Fatal(err)
				return nil, err
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("")
			return nil, err
		}
	}
	sum := sha256.Sum(nil)
	return sum, nil
}
