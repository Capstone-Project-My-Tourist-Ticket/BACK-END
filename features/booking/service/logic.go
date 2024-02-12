package service

import (
	"errors"
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

func (service *bookingService) CancleBooking(userIdLogin int, bookingId string, bookingCore booking.Core) error {
	if bookingCore.Status == "" {
		bookingCore.Status = "cancelled"
	}

	err := service.bookingData.CancleBooking(userIdLogin, bookingId, bookingCore)
	return err
}

// WebhoocksService implements booking.BookingServiceInterface.
func (service *bookingService) WebhoocksService(reqNotif booking.Core) error {
	if reqNotif.ID == "" {
		return errors.New("invalid order id")
	}

	err := service.bookingData.WebhoocksData(reqNotif)
	if err != nil {
		return err
	}

	return nil
}
