package data

import (
	"errors"
	packages "my-tourist-ticket/features/package"

	"gorm.io/gorm"
)

type packageQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) packages.PackageDataInterface {
	return &packageQuery{
		db: db,
	}
}

// Insert implements packages.PackageDataInterface.
func (repo *packageQuery) Insert(benefits []string, input packages.Core) error {
	dataGorm := CoreToModel(input)

	tx := repo.db.Create(&dataGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	lastInsertedID := dataGorm.ID

	for _, value := range benefits {
		benefitValue := Benefit{
			PackageID: lastInsertedID,
			Benefit:   value,
		}

		tb := repo.db.Create(&benefitValue)
		if tb.Error != nil {
			return tb.Error
		}
	}
	return nil
}
