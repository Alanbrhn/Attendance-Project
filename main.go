package main

import (
	"Attendance-Project/database"
	"Attendance-Project/routes"
	"Attendance-Project/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Menghubungkan ke database menggunakan GORM
	database.Connect()

	// Membuat userService menggunakan *gorm.DB
	userService := services.NewUserService(database.DB)

	// Mengonfigurasi router Gin
	r := gin.Default()

	// Menyiapkan routes dengan mengirimkan pointer ke userService
	routes.SetupRoutes(r, userService)

	// Menjalankan server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
