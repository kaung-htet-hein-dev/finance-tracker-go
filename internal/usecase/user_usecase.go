package usecase

import (
	"kaung-htet-hein-dev/finance-tracker-go/internal/domain"
	"kaung-htet-hein-dev/finance-tracker-go/internal/infrastructure/repository"
	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1/request"
	"kaung-htet-hein-dev/finance-tracker-go/pkg"
)

type UserUsecaseInterface interface {
	CreateUser(r *request.CreateUserRequest) (string, error)
}

type userUsecase struct {
	userRepo *repository.UserRepository
}

func NewUserUsecase(userRepo *repository.UserRepository) UserUsecaseInterface {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) CreateUser(r *request.CreateUserRequest) (string, error) {
	token := "sample-token"

	err := u.userRepo.CreateUser(&domain.User{
		FullName: r.FullName,
		Email:    r.Email,
		Password: pkg.HashPassword(r.Password),
	})

	return token, err
}
