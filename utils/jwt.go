package utils

import (
	"Attendance-Project/config"
	"time"
    "errors"

	"github.com/golang-jwt/jwt/v5"

)

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	// Parsing dan memverifikasi token menggunakan secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Pastikan algoritma signing yang digunakan adalah HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		// Kembalikan secret key untuk validasi
		return []byte(config.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// Memastikan token valid (sudah tidak expired, dll)
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
