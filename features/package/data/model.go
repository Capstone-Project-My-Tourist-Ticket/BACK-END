package data

import (
	packages "my-tourist-ticket/features/package"

	"gorm.io/gorm"
)

type Package struct {
	gorm.Model
	TourID      uint
	PackageName string
	Price       int
	JumlahTiket int
	Benefits    []Benefit
}

type Benefit struct {
	gorm.Model
	PackageID uint
	Benefit   string
}

func CoreToModel(input packages.Core) Package {
	return Package{
		TourID:      input.TourID,
		PackageName: input.PackageName,
		Price:       input.Price,
		JumlahTiket: input.JumlahTiket,
	}
}
