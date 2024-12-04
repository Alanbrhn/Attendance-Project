package routes

import (
	"Attendance-Project/controllers"
	"Attendance-Project/services"

	"github.com/gin-gonic/gin"

)

func SetupRoutes(router *gin.Engine, userService *services.UserService) {
	// Initialize UserController
	userController := controllers.NewUserController(userService)

	// Auth routes
	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)

	// User routes
	router.POST("/users/register", userController.Register)
	router.GET("/users/:id", userController.GetUser)
}
