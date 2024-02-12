package data

import (
	"my-tourist-ticket/features/booking"
	pd "my-tourist-ticket/features/package/data"
	td "my-tourist-ticket/features/tour/data"
	ud "my-tourist-ticket/features/user/data"
	vd "my-tourist-ticket/features/voucher/data"
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID string `gorm:"type:varchar(36);primary_key" json:"id"`
	gorm.Model
	UserID      uint  `gorm:"not null"`
	TourID      uint  `gorm:"not null"`
	PackageID   uint  `gorm:"not null"`
	VoucherID   *uint `gorm:"default:null;omitempty"`
	PaymentType string
	GrossAmount int
	Status      string
	VaNumber    string
	Bank        string
	PhoneNumber string
	Greeting    string
	FullName    string
	Email       string
	Quantity    int
	ExpiredAt   time.Time
	User        ud.User
	Tour        td.Tour
	Package     pd.Package
	Voucher     vd.Voucher
}

func CoreToModelBooking(input booking.Core) Booking {
	return Booking{
		ID:          input.ID,
		UserID:      input.UserID,
		TourID:      input.TourID,
		PackageID:   input.PackageID,
		VoucherID:   input.VoucherID,
		PaymentType: input.PaymentType,
		GrossAmount: input.GrossAmount,
		Status:      input.Status,
		VaNumber:    input.VaNumber,
		Bank:        input.Bank,
		PhoneNumber: input.PhoneNumber,
		Greeting:    input.Greeting,
		FullName:    input.FullName,
		Email:       input.Email,
		Quantity:    input.Quantity,
		ExpiredAt:   input.ExpiredAt,
	}
}

func ModelToCoreBooking(model Booking) booking.Core {
	return booking.Core{
		ID:          model.ID,
		UserID:      model.UserID,
		TourID:      model.TourID,
		PackageID:   model.PackageID,
		VoucherID:   model.VoucherID,
		PaymentType: model.PaymentType,
		GrossAmount: model.GrossAmount,
		Status:      model.Status,
		VaNumber:    model.VaNumber,
		Bank:        model.Bank,
		PhoneNumber: model.PhoneNumber,
		Greeting:    model.Greeting,
		FullName:    model.FullName,
		Email:       model.Email,
		Quantity:    model.Quantity,
		ExpiredAt:   model.ExpiredAt,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func CoreToModel(reqNotif booking.Core) Booking {
	return Booking{
		Status: reqNotif.Status,
	}
}

func CoreToModelBookingCancle(input booking.Core) Booking {
	return Booking{
		Status: input.Status,
	}
}
