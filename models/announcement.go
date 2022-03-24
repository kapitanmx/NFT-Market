package models

import (
	"models"
	"time"

	"github.com/google/uuid"
)

type Announcement struct {
	ID             string
	Date           string
	ExpDate        string
	Title          string
	Desc           string
	Imgs           []string
	Price          float64
	AdvertiserName string
	AdvertiserID   string
	Category       string
	Tags           []string
}

func (a *Announcement) CreateAnnoncement(
	ID,
	date time.Time,
	expDate time.Time,
	title,
	desc,
	imgs,
	tags,
	advertiserName,
	advertiserID,
	category string,
	price float64) (*Announcement, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	newDate := SetDate(date)
	newExpDate := SetExpDate(expDate)
	images := SetImgs(imgs)
	tagList := SetTags(tags)
	announcement := &Announcement{
		ID:             id.String(),
		Date:           newDate,
		ExpDate:        newExpDate,
		Title:          title,
		Desc:           desc,
		Imgs:           images,
		AdvertiserName: advertiserName,
		AdvertiserID:   advertiserID,
		Category:       category,
		Price:          price,
		Tags:           tagList,
	}
	if models.IsAnnouncementEmpty() {
		return nil, err
	}
	return announcement, nil
}

func SetDate(t time.Time) string {
	date := t.String()
	return date
}

func SetExpDate(t time.Time) string {
	expDate := t.String()
	return expDate
}

func SetImgs(imgs ...string) []string {
	images := make([]string, 0)
	images = append(images, imgs...)
	return images
}

func SetTags(tags ...string) []string {
	tagList := make([]string, 0)
	tagList = append(tagList, tags...)
	return tagList
}
