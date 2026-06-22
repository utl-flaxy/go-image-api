package services

import (
	"go-image-api/models"
	"go-image-api/repositories"
	"go-image-api/utils"

	"golang.org/x/crypto/bcrypt"
)

func Register(name, email, password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	return repositories.CreateUser(user)
}

func Login(email, password string) (string, error) {

	user, err := repositories.FindByEmail(email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
