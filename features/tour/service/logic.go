package service

import (
	"fmt"
	"mime/multipart"
	"my-tourist-ticket/features/tour"
)

type tourService struct {
	tourData tour.TourDataInterface
}

func NewTour(repo tour.TourDataInterface) tour.TourServiceInterface {
	return &tourService{
		tourData: repo,
	}
}

// GetUserRoleById implements tour.TourServiceInterface.
func (service *tourService) GetUserRoleById(userId int) (string, error) {
	return service.tourData.GetUserRoleById(userId)
}

// Insert implements tour.TourServiceInterface.
func (service *tourService) Insert(userId uint, input tour.Core, image *multipart.FileHeader, thumbnail *multipart.FileHeader) error {
	err := service.tourData.Insert(userId, input, image, thumbnail)
	if err != nil {
		return fmt.Errorf("error creating city: %w", err)
	}

	return nil
}

// Update implements tour.TourServiceInterface.
func (service *tourService) Update(tourId int, input tour.Core, image *multipart.FileHeader, thumbnail *multipart.FileHeader) error {
	err := service.tourData.Update(tourId, input, image, thumbnail)
	if err != nil {
		return fmt.Errorf("error update tour: %w", err)
	}

	return nil
}
