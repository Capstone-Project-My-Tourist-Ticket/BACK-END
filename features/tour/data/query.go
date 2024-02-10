package data

import (
	"errors"
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

func (repo *tourQuery) GetImageByTourId(tourId int) (string, error) {
	var tour Tour
	if err := repo.db.Table("tours").Where("id = ?", tourId).First(&tour).Error; err != nil {
		return "", err
	}

	return tour.Image, nil
}

func (repo *tourQuery) GetThumbnailByTourId(tourId int) (string, error) {
	var tour Tour
	if err := repo.db.Table("tours").Where("id = ?", tourId).First(&tour).Error; err != nil {
		return "", err
	}

	return tour.Thumbnail, nil
}

// Update implements tour.TourDataInterface.
func (repo *tourQuery) Update(tourId int, input tour.Core, image *multipart.FileHeader, thumbnail *multipart.FileHeader) error {
	// Dapatkan image dan thumbnail dari database
	imgGorm, _ := repo.GetImageByTourId(tourId)
	thumbnailGorm, _ := repo.GetThumbnailByTourId(tourId)

	// Hapus image lama jika ada
	if imgGorm != "" {
		publicID := cloudinary.GetPublicID(imgGorm)
		if err := repo.uploadService.Destroy(publicID); err != nil {
			return fmt.Errorf("error destroying old image from Cloudinary: %w", err)
		}
		fmt.Print("image publicID", publicID)
	}

	// Hapus thumbnail lama jika ada
	if thumbnailGorm != "" {
		publicID := cloudinary.GetPublicID(thumbnailGorm)
		if err := repo.uploadService.Destroy(publicID); err != nil {
			return fmt.Errorf("error destroying old thumbnail from Cloudinary: %w", err)
		}
		fmt.Print("thumbnail publicID", publicID)
	}

	// Upload image baru ke Cloudinary
	imageURL, err := repo.uploadService.UploadImage(image)
	if err != nil {
		return fmt.Errorf("error uploading image to Cloudinary: %w", err)
	}

	// Upload thumbnail baru ke Cloudinary
	thumbnailURL, err := repo.uploadService.UploadImage(thumbnail)
	if err != nil {
		return fmt.Errorf("error uploading thumbnail to Cloudinary: %w", err)
	}

	// Perbarui atribut-atribut yang diperlukan
	tourGorm := CoreToModel(input)
	tourGorm.Image = imageURL
	tourGorm.Thumbnail = thumbnailURL

	// Lakukan update data kota di dalam database
	tx := repo.db.Model(&Tour{}).Where("id = ?", tourId).Updates(tourGorm)
	if tx.Error != nil {
		return fmt.Errorf("error updating tour: %w", tx.Error)
	}
	if tx.RowsAffected == 0 {
		return errors.New("error: tour not found")
	}
	return nil
}

// SelectTourById implements tour.TourDataInterface.
func (repo *tourQuery) SelectTourById(tourId int) (tour.Core, error) {
	var tourModel Tour
	if err := repo.db.First(&tourModel, tourId).Error; err != nil {
		return tour.Core{}, err
	}

	return ModelToCore(tourModel), nil
}

// Delete implements tour.TourDataInterface.
func (repo *tourQuery) Delete(tourId int) error {
	dataTour, _ := repo.SelectTourById(tourId)

	if dataTour.ID != uint(tourId) {
		return errors.New("tour not found")
	}

	tx := repo.db.Where("id = ?", tourId).Delete(&Tour{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("error not found")
	}
	return nil
}

// SelectAllTour implements tour.TourDataInterface.
func (repo *tourQuery) SelectAllTour(page int, limit int) ([]tour.Core, int, error) {
	var tours []Tour
	query := repo.db.Order("created_at desc")

	var totalData int64
	err := query.Model(&tours).Count(&totalData).Error
	if err != nil {
		return nil, 0, err
	}

	totalPage := int((totalData + int64(limit) - 1) / int64(limit))

	err = query.Limit(limit).Offset((page - 1) * limit).Find(&tours).Error
	if err != nil {
		return nil, 0, err
	}

	var tourCores []tour.Core
	for _, t := range tours {
		tourCores = append(tourCores, ModelToCore(t))
	}

	return tourCores, totalPage, nil
}

// SelectTourByPengelola implements tour.TourDataInterface.
func (repo *tourQuery) SelectTourByPengelola(userId int, page, limit int) ([]tour.Core, int, error) {
	var tourDataGorms []Tour
	query := repo.db.Where("user_id = ?", userId)

	var totalData int64
	err := query.Model(&tourDataGorms).Count(&totalData).Error
	if err != nil {
		return nil, 0, err
	}

	totalPage := int((totalData + int64(limit) - 1) / int64(limit))

	err = query.Limit(limit).Offset((page - 1) * limit).Find(&tourDataGorms).Error
	if err != nil {
		return nil, 0, err
	}

	var results []tour.Core
	for _, tourDataGorm := range tourDataGorms {
		result := ModelToCore(tourDataGorm)
		results = append(results, result)
	}

	return results, totalPage, nil
}
