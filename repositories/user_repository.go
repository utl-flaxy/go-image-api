package repositories

import (
	"go-image-api/config"
	"go-image-api/models"
)

func CreateUser(user models.User) error {
	return config.DB.Create(&user).Error
}

func FindByEmail(email string) (models.User, error) {

	var user models.User

	err := config.DB.
		Where("email = ?", email).
		First(&user).
		Error

	return user, err
}
