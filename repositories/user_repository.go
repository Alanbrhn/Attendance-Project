package repositories

import (
	"Attendance-Project/models"
	"errors"

	"gorm.io/gorm"
)

// UserRepository struct
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser creates a new user in the database
func (ur *UserRepository) CreateUser(user models.User) (models.User, error) {
	// GORM query to insert user into the database
	if err := ur.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// GetUserByID fetches a user by ID
func (ur *UserRepository) GetUserByID(id string) (models.User, error) {
	var user models.User
	// GORM query to fetch user by ID
	if err := ur.DB.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}
	return user, nil
}