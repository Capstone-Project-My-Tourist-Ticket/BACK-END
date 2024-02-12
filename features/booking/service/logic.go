package service

import (
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
