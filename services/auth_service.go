package services

import (
	"errors"
	"gin-quickstart/config"
	"gin-quickstart/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var JWTSecret = []byte(os.Getenv("JWT_SECRET"))

func SignUp(user *models.User) (*models.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashed)
	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func Login(email, password string) (string, *models.User, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", nil, errors.New("user not found")
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, errors.New("invalid password")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", nil, err
	}

	return tokenString, &user, nil
}
