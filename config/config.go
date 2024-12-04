package config

import (
	"log"
	"os"
)

var (
	JWTSecret    string
	DatabaseDSN  string
)

func LoadConfig() {
	JWTSecret = os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	DatabaseDSN = os.Getenv("DATABASE_DSN")
	if DatabaseDSN == "" {
		log.Fatal("DATABASE_DSN is required")
	}
}
