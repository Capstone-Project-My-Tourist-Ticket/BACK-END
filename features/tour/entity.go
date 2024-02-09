package tour

import "mime/multipart"

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
}

type TourDataInterface interface {
	GetUserRoleById(userId int) (string, error)
	Insert(userId uint, input Core, image *multipart.FileHeader, thumbnail *multipart.FileHeader) error
	Update(tourId int, input Core, image *multipart.FileHeader, thumbnail *multipart.FileHeader) error
}

type TourServiceInterface interface {
	GetUserRoleById(userId int) (string, error)
	Insert(userId uint, input Core, image *multipart.FileHeader, thumbnail *multipart.FileHeader) error
	Update(tourId int, input Core, image *multipart.FileHeader, thumbnail *multipart.FileHeader) error
}
