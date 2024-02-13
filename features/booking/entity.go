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

type ReviewCore struct {
	ID         uint
	BookingID  string
	UserID     uint
	TextReview string
	StartRate  float64
	Booking    Core
	User       user.Core
}

// interface untuk Data Layer
type BookingDataInterface interface {
	InsertBooking(userIdLogin int, inputBooking Core) (*Core, error)
	CancleBooking(userIdLogin int, orderId string, bookingCore Core) error
	InsertBookingReview(inputReview ReviewCore) error
	WebhoocksData(reqNotif Core) error
	SelectBookingUser(userIdLogin int) ([]Core, error)
	SelectAllBooking(page, limit int) ([]Core, int, error)
	GetUserRoleById(userId int) (string, error)
	SelectAllBookingPengelola(pengelolaID int, page, limit int) ([]Core, int, error)
}

// interface untuk Service Layer
type BookingServiceInterface interface {
	CreateBooking(userIdLogin int, inputBooking Core) (*Core, error)
	CancleBooking(userIdLogin int, orderId string, bookingCore Core) error
	CreateBookingReview(inputReview ReviewCore) error
	WebhoocksService(reqNotif Core) error
	GetBookingUser(userIdLogin int) ([]Core, error)
	SelectAllBooking(page, limit int) ([]Core, int, error)
	GetUserRoleById(userId int) (string, error)
	SelectAllBookingPengelola(pengelolaID int, page, limit int) ([]Core, int, error)
}
