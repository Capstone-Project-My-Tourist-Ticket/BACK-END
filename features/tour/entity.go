package tour

import (
	"mime/multipart"
	"time"
)

type Core struct {
	ID          uint
	CityId      uint
	UserId      uint
	TourName    string
	Description string
	Image       string
	Thumbnail   string
	Address     string
	Latitude    float64
	Longitude   float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TourDataInterface interface {
	GetUserRoleById(userId int) (string, error)
	Insert(userId uint, input Core, image *multipart.FileHeader, thumbnail *multipart.FileHeader) error
	Update(tourId int, input Core, image *multipart.FileHeader, thumbnail *multipart.FileHeader) error
	SelectTourById(tourId int) (Core, error)
	Delete(tourId int) error
	SelectAllTour(page, limit int) ([]Core, int, error)
	SelectTourByPengelola(userId int, page, limit int) ([]Core, int, error)
	GetTourByCityID(cityID uint, page, limit int) ([]Core, int, error)
}

type TourServiceInterface interface {
	GetUserRoleById(userId int) (string, error)
	Insert(userId uint, input Core, image *multipart.FileHeader, thumbnail *multipart.FileHeader) error
	Update(tourId int, input Core, image *multipart.FileHeader, thumbnail *multipart.FileHeader) error
	SelectTourById(tourId int) (Core, error)
	Delete(tourId int) error
	SelectAllTour(page, limit int) ([]Core, int, error)
	SelectTourByPengelola(userId int, page, limit int) ([]Core, int, error)
	GetTourByCityID(cityID uint, page, limit int) ([]Core, int, error)
}
