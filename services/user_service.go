package services

import (
	"Attendance-Project/models"
	"Attendance-Project/repositories"

	"gorm.io/gorm"
)

// UserService struct
type UserService struct {
	UserRepository repositories.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(db *gorm.DB) *UserService {
    return &UserService{
        UserRepository: *repositories.NewUserRepository(db),
    }
}

// CreateUser creates a new user
func (us *UserService) CreateUser(user models.User) (models.User, error) {
	// Add validation and any additional logic as needed
	return us.UserRepository.CreateUser(user)
}

// GetUserByID fetches a user by ID
func (us *UserService) GetUserByID(id string) (models.User, error) {
	user, err := us.UserRepository.GetUserByID(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
