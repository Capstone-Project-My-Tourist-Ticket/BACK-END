package data

import (
	"my-tourist-ticket/features/city"
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

type Report struct {
	gorm.Model
	UserId     uint
	TourId     uint
	TextReport string
	User       _userData.User
	Tour       Tour
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
		City: city.Core{
			ID:       t.City.ID,
			CityName: t.City.City,
			// Description: t.City.Description,
			Image:     t.City.Image,
			Thumbnail: t.City.Thumbnail,
		},
	}
}

func (t Tour) ModelToCoreTourBooking() tour.Core {
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

func ModelToCoreList(data []Tour) []tour.Core {
	var results []tour.Core
	for _, t := range data {
		results = append(results, ModelToCore(t))
	}
	return results
}

func CoreReportToModelReport(tr tour.ReportCore) Report {
	return Report{
		TourId:     tr.TourId,
		UserId:     tr.UserId,
		TextReport: tr.TextReport,
	}
}

func ModelToReportCore(r Report) tour.ReportCore {
	return tour.ReportCore{
		ID:         r.ID,
		UserId:     r.UserId,
		TourId:     r.TourId,
		TextReport: r.TextReport,
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}
