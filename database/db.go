package database

import (
	"Attendance-Project/config"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect menghubungkan aplikasi dengan database SQL Server
func Connect() {
	var err error

	// Gunakan sqlserver.Open() untuk koneksi ke SQL Server
	DB, err = gorm.Open(sqlserver.Open(config.DatabaseDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	log.Println("Successfully connected to the database")
}
