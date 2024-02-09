package data

import (
	"errors"
	"fmt"
	"mime/multipart"
	"my-tourist-ticket/features/city"
	"my-tourist-ticket/features/user"
	"my-tourist-ticket/utils/cloudinary"

	"gorm.io/gorm"
)

type cityQuery struct {
	db            *gorm.DB
	uploadService cloudinary.CloudinaryUploaderInterface
}

func NewCity(db *gorm.DB, cloud cloudinary.CloudinaryUploaderInterface) city.CityDataInterface {
	return &cityQuery{
		db:            db,
		uploadService: cloud,
	}
}

// GetUserRoleById
func (repo *cityQuery) GetUserRoleById(userId int) (string, error) {
	var user user.Core
	if err := repo.db.Table("users").Where("id = ?", userId).First(&user).Error; err != nil {
		return "", err
	}

	return user.Role, nil
}

// Insert implements city.CityDataInterface.
func (repo *cityQuery) Insert(input city.Core, image *multipart.FileHeader, thumbnail *multipart.FileHeader) error {
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
	newCity := CoreToModel(input)
	newCity.Image = imageURL
	newCity.Thumbnail = thumbnailURL

	if err := repo.db.Create(&newCity).Error; err != nil {
		return fmt.Errorf("error inserting city: %w", err)
	}

	return nil
}

func (repo *cityQuery) GetImageByCityId(cityId int) (string, error) {
	var city City
	if err := repo.db.Table("cities").Where("id = ?", cityId).First(&city).Error; err != nil {
		return "", err
	}

	return city.Image, nil
}

func (repo *cityQuery) GetThumbnailByCityId(cityId int) (string, error) {
	var city City
	if err := repo.db.Table("cities").Where("id = ?", cityId).First(&city).Error; err != nil {
		return "", err
	}

	return city.Thumbnail, nil
}

// Update implements city.CityDataInterface.
func (repo *cityQuery) Update(cityId int, input city.Core, image *multipart.FileHeader, thumbnail *multipart.FileHeader) error {
	// Dapatkan image dan thumbnail dari database
	imgGorm, _ := repo.GetImageByCityId(cityId)
	thumbnailGorm, _ := repo.GetThumbnailByCityId(cityId)

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
	cityGorm := CoreToModel(input)
	cityGorm.Image = imageURL
	cityGorm.Thumbnail = thumbnailURL

	// Lakukan update data kota di dalam database
	tx := repo.db.Model(&City{}).Where("id = ?", cityId).Updates(cityGorm)
	if tx.Error != nil {
		return fmt.Errorf("error updating city: %w", tx.Error)
	}
	if tx.RowsAffected == 0 {
		return errors.New("error: city not found")
	}
	return nil
}

// Delete implements city.CityDataInterface.
func (*cityQuery) Delete(cityId int) error {
	panic("unimplemented")
}

// SelectCityById implements city.CityDataInterface.
func (repo *cityQuery) SelectCityById(cityId int) (city.Core, error) {
	var cityModel City
	if err := repo.db.First(&cityModel, cityId).Error; err != nil {
		return city.Core{}, err
	}

	return ModelToCore(cityModel), nil
}
