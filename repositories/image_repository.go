package repositories

import (
	"go-image-api/config"
	"go-image-api/models"
)

func CreateImage(image models.Image) error {

	return config.DB.Create(&image).Error
}

func GetImagesByUserID(userID uint) ([]models.Image, error) {

	var images []models.Image

	err := config.DB.
		Where("user_id = ?", userID).
		Find(&images).
		Error

	return images, err
}

func FindImageByID(id string) (models.Image, error) {

	var image models.Image

	err := config.DB.First(&image, id).Error

	return image, err
}

func DeleteImage(id string) error {

	return config.DB.Delete(&models.Image{}, id).Error
}
