package repositories

import (
	"gin-quickstart/config"
	"gin-quickstart/models"
)

func CreateUser(user *models.User) error {
    return config.DB.Create(user).Error
}

func GetUsers() ([]models.User, error) {
    var users []models.User
    err := config.DB.Find(&users).Error
    return users, err
}
