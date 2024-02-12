package booking

import (
	packages "my-tourist-ticket/features/package"
	"my-tourist-ticket/features/tour"
	"my-tourist-ticket/features/user"
	"my-tourist-ticket/features/voucher"
	"time"
)

type Core struct {
	ID          string
	UserID      uint
	TourID      uint
	PackageID   uint
	VoucherID   *uint
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
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        user.Core
	Tour        tour.Core
	Package     packages.Core
	Voucher     voucher.Core
}

// interface untuk Data Layer
type BookingDataInterface interface {
	InsertBooking(userIdLogin int, inputBooking Core) (*Core, error)
}

// interface untuk Service Layer
type BookingServiceInterface interface {
	CreateBooking(userIdLogin int, inputBooking Core) (*Core, error)
}
