package data

import (
	"errors"
	"my-tourist-ticket/features/voucher"

	"gorm.io/gorm"
)

type voucherQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) voucher.VoucherDataInterface {
	return &voucherQuery{
		db: db,
	}
}

// Insert implements voucher.VoucherDataInterface.
func (repo *voucherQuery) Insert(input voucher.Core) error {
	dataGorm := CoreToModel(input)

	tx := repo.db.Create(&dataGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}
