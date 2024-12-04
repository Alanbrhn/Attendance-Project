package controllers

import (
	"Attendance-Project/services"
	"Attendance-Project/models"
	"net/http"

	"github.com/gin-gonic/gin"

)

// UserController struct
type UserController struct {
	UserService *services.UserService
}

// NewUserController creates a new UserController
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// Register is the handler for user registration
func (uc *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := uc.UserService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdUser)
}

// GetUser is the handler for getting a user by ID
func (uc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.UserService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
