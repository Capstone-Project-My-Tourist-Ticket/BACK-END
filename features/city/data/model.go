package data

import (
	"my-tourist-ticket/features/city"

	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	City        string `gorm:"unique"`
	Description string
	Image       string
	Thumbnail   string
}

func CoreToModel(c city.Core) City {
	return City{
		City:        c.CityName,
		Description: c.Description,
		Image:       c.Image,
		Thumbnail:   c.Thumbnail,
	}
}
