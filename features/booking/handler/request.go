package handler

import (
	"my-tourist-ticket/features/booking"

	"github.com/google/uuid"
)

type BookingRequest struct {
	TourID      uint   `json:"tour_id" form:"tour_id"`
	PackageID   uint   `json:"package_id" form:"package_id"`
	VoucherID   uint   `json:"voucher_id" form:"voucher_id"`
	PaymentType string `json:"payment_type" form:"payment_type"`
	GrossAmount int    `json:"gross_amount" form:"gross_amount"`
	Bank        string `json:"bank" form:"bank"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Greeting    string `json:"greeting" form:"greeting"`
	FullName    string `json:"full_name" form:"full_name"`
	Email       string `json:"email" form:"email"`
	Quantity    int    `json:"quantity" form:"quantity"`
}

type WebhoocksRequest struct {
	BookingID         string `json:"order_id"`
	TransactionStatus string `json:"transaction_status"`
}

func RequestToCoreBooking(input BookingRequest, userIdLogin uint) booking.Core {
	return booking.Core{
		ID:          uuid.New().String(),
		UserID:      userIdLogin,
		TourID:      input.TourID,
		PackageID:   input.PackageID,
		VoucherID:   &input.VoucherID,
		PaymentType: input.PaymentType,
		GrossAmount: input.GrossAmount,
		Bank:        input.Bank,
		PhoneNumber: input.PhoneNumber,
		Greeting:    input.Greeting,
		FullName:    input.FullName,
		Email:       input.Email,
		Quantity:    input.Quantity,
	}
}

func WebhoocksRequestToCore(input WebhoocksRequest) booking.Core {
	return booking.Core{
		ID:     input.BookingID,
		Status: input.TransactionStatus,
	}
}
