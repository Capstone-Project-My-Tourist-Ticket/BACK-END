package handler

import (
	"my-tourist-ticket/features/tour"
)

type TourResponse struct {
	ID          uint         `json:"id"`
	CityId      uint         `json:"city_id"`
	UserId      uint         `json:"user_id"`
	TourName    string       `json:"tour_name"`
	Description string       `json:"description"`
	Image       string       `json:"image"`
	Thumbnail   string       `json:"thumbnail"`
	Address     string       `json:"address"`
	Latitude    float64      `json:"latitude"`
	Longitude   float64      `json:"longitude"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
	City        CityResponse `json:"city"`
}

type ReportResponse struct {
	ID         uint   `json:"id"`
	TourId     uint   `json:"tour_id"`
	UserId     uint   `json:"user_id"`
	TextReport string `json:"text_report"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type CityResponse struct {
	ID       uint   `json:"id"`
	CityName string `json:"city_name"`
	// Description string `json:"description"`
	Image     string `json:"image"`
	Thumbnail string `json:"thumbnail"`
}

func ModelToResponse(tourModel tour.Core) TourResponse {
	return TourResponse{
		ID:          tourModel.ID,
		CityId:      tourModel.CityId,
		UserId:      tourModel.UserId,
		TourName:    tourModel.TourName,
		Description: tourModel.Description,
		Image:       tourModel.Image,
		Thumbnail:   tourModel.Thumbnail,
		Address:     tourModel.Address,
		Latitude:    tourModel.Latitude,
		Longitude:   tourModel.Longitude,
		CreatedAt:   tourModel.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   tourModel.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func CoreToGetAllResponseTour(data tour.Core) TourResponse {
	return TourResponse{
		ID:          data.ID,
		CityId:      data.CityId,
		UserId:      data.UserId,
		TourName:    data.TourName,
		Description: data.Description,
		Image:       data.Image,
		Thumbnail:   data.Thumbnail,
		Address:     data.Address,
		Latitude:    data.Latitude,
		Longitude:   data.Longitude,
		CreatedAt:   data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   data.UpdatedAt.Format("2006-01-02 15:04:05"),
		City: CityResponse{
			ID:       data.City.ID,
			CityName: data.City.CityName,
			// Description: data.City.Description,
			Image:     data.City.Image,
			Thumbnail: data.City.Thumbnail,
		},
	}
}

func CoreToResponseListGetAllTour(data []tour.Core) []TourResponse {
	var results []TourResponse
	for _, v := range data {
		results = append(results, CoreToGetAllResponseTour(v))
	}
	return results
}

func ModelReportToReportResponse(reportModel tour.ReportCore) ReportResponse {
	return ReportResponse{
		ID:         reportModel.ID,
		TourId:     reportModel.TourId,
		UserId:     reportModel.UserId,
		TextReport: reportModel.TextReport,
		CreatedAt:  reportModel.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  reportModel.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func CoreReportToResponseListGetReportTour(data []tour.ReportCore) []ReportResponse {
	var results []ReportResponse
	for _, v := range data {
		results = append(results, ModelReportToReportResponse(v))
	}
	return results
}
