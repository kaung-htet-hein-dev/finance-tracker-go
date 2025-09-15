package usecase

import "kaung-htet-hein-dev/finance-tracker-go/internal/infrastructure/repository"

type UserUsecase interface {
	Login() error
}

type userUsecase struct {
	userRepo *repository.UserRepository
}

func NewUserUsecase(userRepo *repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) Login() error {
	return nil
}
