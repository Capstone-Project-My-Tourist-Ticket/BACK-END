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

// GetUserRoleById implements booking.BookingServiceInterface.
func (service *bookingService) GetUserRoleById(userId int) (string, error) {
	return service.bookingData.GetUserRoleById(userId)
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

func (service *bookingService) GetBookingUser(userIdLogin int) ([]booking.Core, error) {
	result, err := service.bookingData.SelectBookingUser(userIdLogin)
	return result, err
}

// SelectAllBooking implements booking.BookingServiceInterface.
func (service *bookingService) SelectAllBooking(page int, limit int) ([]booking.Core, int, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 8
	}

	bookings, totalPage, err := service.bookingData.SelectAllBooking(page, limit)
	if err != nil {
		return nil, 0, err
	}

	return bookings, totalPage, nil
}

// SelectAllPengelola implements booking.BookingServiceInterface.
func (service *bookingService) SelectAllBookingPengelola(pengelolaID int, page int, limit int) ([]booking.Core, int, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 8
	}

	bookings, totalPage, err := service.bookingData.SelectAllBookingPengelola(pengelolaID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	return bookings, totalPage, nil
}
