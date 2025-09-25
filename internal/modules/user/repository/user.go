package repository

import (
	"github.com/solv1k/croco-api/internal/modules/user/models"
	"gorm.io/gorm"
)

// User GORM repository
type UserRepository struct {
	DB *gorm.DB
}

// Repository constructor
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Retrieve user by ID
func (r *UserRepository) GetByID(userID string) (models.User, error) {
	var user models.User
	err := r.DB.Table("users").Where("id = ?", userID).First(&user).Error
	return user, err
}

// Retrieve first user by email and create if not exists
func (r *UserRepository) FirstOrCreate(user *models.User) error {
	return r.DB.Where("email = ?", user.Email).FirstOrCreate(&user).Error
}
