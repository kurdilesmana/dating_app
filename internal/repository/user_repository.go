package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kurdilesmana/dating_app/internal/domain"
)

type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// CreateUser creates a new user in the database.
func (ur *UserRepository) Create(user *domain.User) error {
	if err := ur.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// FindByUsername in the database.
func (ur *UserRepository) FindByUsername(email string) (*domain.User, error) {
	return nil, nil
}
