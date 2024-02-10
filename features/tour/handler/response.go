package handler

import "my-tourist-ticket/features/tour"

type TourResponse struct {
	ID          uint    `json:"id"`
	CityId      uint    `json:"city_id"`
	UserId      uint    `json:"user_id"`
	TourName    string  `json:"tour_name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Thumbnail   string  `json:"thumbnail"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
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
