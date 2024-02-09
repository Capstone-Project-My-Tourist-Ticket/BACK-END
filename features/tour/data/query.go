package data

import (
	"fmt"
	"mime/multipart"
	"my-tourist-ticket/features/tour"
	"my-tourist-ticket/features/user"
	"my-tourist-ticket/utils/cloudinary"

	"gorm.io/gorm"
)

type tourQuery struct {
	db            *gorm.DB
	uploadService cloudinary.CloudinaryUploaderInterface
}

// Insert implements tour.TourDataInterface.

func NewTour(db *gorm.DB, cloud cloudinary.CloudinaryUploaderInterface) tour.TourDataInterface {
	return &tourQuery{
		db:            db,
		uploadService: cloud,
	}
}

// GetUserRoleById implements tour.TourDataInterface.
func (repo *tourQuery) GetUserRoleById(userId int) (string, error) {
	var user user.Core
	if err := repo.db.Table("users").Where("id = ?", userId).First(&user).Error; err != nil {
		return "", err
	}

	return user.Role, nil
}

func (repo *tourQuery) Insert(userId uint, input tour.Core, image *multipart.FileHeader, thumbnail *multipart.FileHeader) error {
	// Upload image dan thumbnail ke Cloudinary
	imageURL, err := repo.uploadService.UploadImage(image)
	if err != nil {
		return fmt.Errorf("error uploading image to Cloudinary: %w", err)
	}

	thumbnailURL, err := repo.uploadService.UploadImage(thumbnail)
	if err != nil {
		return fmt.Errorf("error uploading thumbnail to Cloudinary: %w", err)
	}

	// Buat objek City dengan URL gambar dan thumbnail yang telah di-upload
	newTour := CoreToModel(input)
	newTour.Image = imageURL
	newTour.Thumbnail = thumbnailURL

	if err := repo.db.Create(&newTour).Error; err != nil {
		return fmt.Errorf("error inserting city: %w", err)
	}

	return nil
}
