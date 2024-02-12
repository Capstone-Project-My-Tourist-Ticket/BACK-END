package handler

import (
	"my-tourist-ticket/features/booking"
	"time"
)

type BookingResponse struct {
	ID          string    `json:"booking_id"`
	UserID      uint      `json:"user_id"`
	TourID      uint      `json:"tour_id"`
	PackageID   uint      `json:"package_id"`
	VoucherID   *uint     `json:"voucher_id"`
	PaymentType string    `json:"payment_type"`
	GrossAmount int       `json:"gross_amount"`
	Status      string    `json:"status"`
	VaNumber    string    `json:"va_number"`
	Bank        string    `json:"bank"`
	PhoneNumber string    `json:"phone_number"`
	Greeting    string    `json:"greeting"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	Quantity    int       `json:"quantity"`
	ExpiredAt   time.Time `json:"payment_expired"`
	CreatedAt   time.Time `json:"created_at"`
}

func CoreToResponseBooking(core *booking.Core) BookingResponse {
	return BookingResponse{
		ID:          core.ID,
		UserID:      core.UserID,
		TourID:      core.TourID,
		PackageID:   core.PackageID,
		VoucherID:   core.VoucherID,
		PaymentType: core.PaymentType,
		GrossAmount: core.GrossAmount,
		Status:      core.Status,
		VaNumber:    core.VaNumber,
		Bank:        core.Bank,
		PhoneNumber: core.PhoneNumber,
		Greeting:    core.Greeting,
		FullName:    core.FullName,
		Email:       core.Email,
		Quantity:    core.Quantity,
		ExpiredAt:   core.ExpiredAt,
		CreatedAt:   core.CreatedAt,
	}
}
