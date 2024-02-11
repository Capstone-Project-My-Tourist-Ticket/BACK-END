package voucher

import "time"

type Core struct {
	ID             uint
	Name           string
	Code           string
	Description    string
	DiscountValue  int
	ExpiredVoucher string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// interface untuk Data Layer
type VoucherDataInterface interface {
	Insert(input Core) error
	SelectAllVoucher() ([]Core, error)
	Update(voucherId int, input Core) error
}

// interface untuk Service Layer
type VoucherServiceInterface interface {
	Create(input Core) error
	SelectAllVoucher() ([]Core, error)
	Update(voucherId int, input Core) error
}
