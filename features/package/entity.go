package packages

import "time"

type Core struct {
	ID          uint
	TourID      uint
	PackageName string
	Price       int
	JumlahTiket int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Benefits    []BenefitCore
}

type BenefitCore struct {
	ID        uint
	PackageID uint
	Benefit   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// interface untuk Data Layer
type PackageDataInterface interface {
	Insert(benefits []string, input Core) error
	SelectByTourId(tourId uint) ([]Core, error)
}

// interface untuk Service Layer
type PackageServiceInterface interface {
	Create(benefits []string, input Core) error
	GetByTourId(tourId uint) ([]Core, error)
}
