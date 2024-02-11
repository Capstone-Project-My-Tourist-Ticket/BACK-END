package data

import (
	"my-tourist-ticket/features/voucher"
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	gorm.Model
	Name           string `gorm:"not null;unique"`
	Code           string `gorm:"not null;unique"`
	Description    string
	DiscountValue  int `gorm:"not null"`
	ExpiredVoucher time.Time
}

func CoreToModel(input voucher.Core) Voucher {
	expiredTimeVoucher, _ := time.Parse("2006-01-02", input.ExpiredVoucher)
	return Voucher{
		Name:           input.Name,
		Code:           input.Code,
		Description:    input.Description,
		DiscountValue:  input.DiscountValue,
		ExpiredVoucher: expiredTimeVoucher,
	}
}
