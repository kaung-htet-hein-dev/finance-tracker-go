package repository

import (
	"errors"
	"kaung-htet-hein-dev/finance-tracker-go/internal/domain"
	"kaung-htet-hein-dev/finance-tracker-go/pkg"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		return pkg.HandleGormError(err, "user")
	}

	return nil
}

func (r *UserRepository) LoginUser(email, password string) (*domain.User, error) {
	var user domain.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, pkg.HandleGormError(err, "user")
	}

	if !pkg.ComparePassword(user.Password, password) {
		return nil, errors.New("invalid credentials")
	}

	if !user.IsConfirmed {
		return nil, errors.New("please confirm your email before logging in")
	}

	return &user, nil
}
