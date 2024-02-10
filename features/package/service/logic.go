package service

import (
	packages "my-tourist-ticket/features/package"
)

type packageService struct {
	packageData packages.PackageDataInterface
}

func New(repo packages.PackageDataInterface) packages.PackageServiceInterface {
	return &packageService{
		packageData: repo,
	}
}

// Create implements packages.PackageServiceInterface.
func (service *packageService) Create(benefits []string, input packages.Core) error {
	if input.JumlahTiket == 0 {
		input.JumlahTiket = 1
	}

	err := service.packageData.Insert(benefits, input)
	if err != nil {
		return err
	}

	return nil
}

// GetByTourId implements packages.PackageServiceInterface.
func (service *packageService) GetByTourId(tourId uint) ([]packages.Core, error) {
	packages, err := service.packageData.SelectByTourId(tourId)
	if err != nil {
		return nil, err
	}
	return packages, nil
}
