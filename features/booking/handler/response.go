package handler

import (
	"my-tourist-ticket/features/booking"
)

type BookingResponse struct {
	ID          string          `json:"booking_id"`
	UserID      uint            `json:"user_id"`
	TourID      uint            `json:"tour_id"`
	PackageID   uint            `json:"package_id"`
	VoucherID   *uint           `json:"voucher_id"`
	PaymentType string          `json:"payment_type"`
	GrossAmount int             `json:"gross_amount"`
	Status      string          `json:"status"`
	VaNumber    string          `json:"va_number"`
	Bank        string          `json:"bank"`
	PhoneNumber string          `json:"phone_number"`
	Greeting    string          `json:"greeting"`
	FullName    string          `json:"full_name"`
	Email       string          `json:"email"`
	Quantity    int             `json:"quantity"`
	ExpiredAt   string          `json:"payment_expired"`
	CreatedAt   string          `json:"created_at"`
	Tour        TourResponse    `json:"tour"`
	Package     PackageResponse `json:"package"`
}

type TourResponse struct {
	ID       uint   `json:"id"`
	TourName string `json:"tour_name"`
}

type PackageResponse struct {
	ID          uint   `json:"id"`
	PackageName string `json:"package_name"`
	Price       int    `json:"price"`
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
		ExpiredAt:   core.ExpiredAt.Format("2006-01-02 15:04:05"),
		CreatedAt:   core.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func CoreToResponse(b booking.Core) BookingResponse {
	return BookingResponse{
		ID:          b.ID,
		UserID:      b.UserID,
		TourID:      b.TourID,
		PackageID:   b.PackageID,
		VoucherID:   b.VoucherID,
		PaymentType: b.PaymentType,
		GrossAmount: b.GrossAmount,
		Status:      b.Status,
		VaNumber:    b.VaNumber,
		Bank:        b.Bank,
		PhoneNumber: b.PhoneNumber,
		Greeting:    b.Greeting,
		FullName:    b.FullName,
		Email:       b.Email,
		Quantity:    b.Quantity,
		ExpiredAt:   b.ExpiredAt.Format("2006-01-02 15:04:05"),
		CreatedAt:   b.CreatedAt.Format("2006-01-02 15:04:05"),
		Tour: TourResponse{
			ID:       b.Tour.ID,
			TourName: b.Tour.TourName,
		},
		Package: PackageResponse{
			ID:          b.Package.ID,
			PackageName: b.Package.PackageName,
			Price:       b.Package.Price,
		},
	}
}

func CoreToResponseList(p []booking.Core) []BookingResponse {
	var results []BookingResponse
	for _, v := range p {
		results = append(results, CoreToResponse(v))
	}
	return results
}
