package services

import (
	"errors"
	"gin-quickstart/config"
	"gin-quickstart/models"
	"gin-quickstart/repositories"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)

	return repositories.CreateUser(user)
}
func GetUsers() ([]models.User, error) {
	return repositories.GetUsers()
}

func GetUserById(id uint) (*models.User, error) {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
