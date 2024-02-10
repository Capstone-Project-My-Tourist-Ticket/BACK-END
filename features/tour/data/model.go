package data

import (
	_cityData "my-tourist-ticket/features/city/data"
	"my-tourist-ticket/features/tour"
	_userData "my-tourist-ticket/features/user/data"

	"gorm.io/gorm"
)

type Tour struct {
	gorm.Model
	CityId      uint
	UserId      uint
	TourName    string
	Description string
	Image       string
	Thumbnail   string
	Addres      string
	Latitude    float64 `gorm:"type:decimal(10,8)"`
	Longitude   float64 `gorm:"type:decimal(11,8)"`
	User        _userData.User
	City        _cityData.City
}

func CoreToModel(t tour.Core) Tour {
	return Tour{
		CityId:      t.CityId,
		UserId:      t.UserId,
		TourName:    t.TourName,
		Description: t.Description,
		Image:       t.Image,
		Thumbnail:   t.Thumbnail,
		Addres:      t.Address,
		Latitude:    t.Latitude,
		Longitude:   t.Longitude,
	}
}

func ModelToCore(t Tour) tour.Core {
	return tour.Core{
		ID:          t.ID,
		CityId:      t.CityId,
		UserId:      t.UserId,
		TourName:    t.TourName,
		Description: t.Description,
		Image:       t.Image,
		Thumbnail:   t.Thumbnail,
		Address:     t.Addres,
		Latitude:    t.Latitude,
		Longitude:   t.Longitude,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
