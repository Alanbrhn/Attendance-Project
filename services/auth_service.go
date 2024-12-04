package services

import (
	"Attendance-Project/database"
	"Attendance-Project/models"
	"Attendance-Project/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(user models.User) error {
	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Simpan ke database
	if err := database.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func LoginUser(credentials models.UserCredentials) (string, error) {
	var user models.User
	if err := database.DB.Where("username = ?", credentials.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("user not found")
		}
		return "", err
	}

	// Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate token JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
