package checkers

import (
	"models"
	"regexp"
)

func (u *models.User) IsUserEmpty() bool {
	userMap := structs.Map(u)
	for _, i range userMap {
		if i != "" {
			return false
		}
	}
	return true
}

func (a *models.Announcement) IsAnnouncementEmpty() bool {
	aMap := structs.Map(a)
	for _, i range aMap {
		if i != "" {
			return false
		}
	}
	return true
}



