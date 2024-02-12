package service

import (
	"errors"
	"fmt"
	"my-tourist-ticket/features/booking"
)

type bookingService struct {
	bookingData booking.BookingDataInterface
}

func New(repo booking.BookingDataInterface) booking.BookingServiceInterface {
	return &bookingService{
		bookingData: repo,
	}
}

// CreateBooking implements booking.BookingServiceInterface.
func (service *bookingService) CreateBooking(userIdLogin int, inputBooking booking.Core) (*booking.Core, error) {
	payment, err := service.bookingData.InsertBooking(userIdLogin, inputBooking)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

// CreateBookingReview implements booking.BookingServiceInterface.
func (service *bookingService) CreateBookingReview(inputReview booking.ReviewCore) error {
	if inputReview.TextReview == "" {
		return errors.New("text review is required")
	}
	if inputReview.StartRate == 0 {
		return errors.New("rate is required")
	}

	err := service.bookingData.InsertBookingReview(inputReview)
	if err != nil {
		return fmt.Errorf("error creating review: %w", err)
	}

	return nil
}
