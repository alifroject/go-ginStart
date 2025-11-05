package services

import (
	"gin-quickstart/models"
	"gin-quickstart/repositories"
)

func CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}
func GetUsers() ([]models.User, error) {
	return repositories.GetUsers()
}
