package usecase

import (
	"kaung-htet-hein-dev/finance-tracker-go/internal/domain"
	"kaung-htet-hein-dev/finance-tracker-go/internal/infrastructure/auth"
	"kaung-htet-hein-dev/finance-tracker-go/internal/infrastructure/repository"
	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1/request"
	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1/response"
	"kaung-htet-hein-dev/finance-tracker-go/pkg"
)

type UserUsecaseInterface interface {
	CreateUser(r *request.CreateUserRequest) (*response.AuthResponse, error)
	LoginUser(r *request.LoginUserRequest) (*response.AuthResponse, error)
	GetUserByID(userID uint) (*domain.User, error)
}

type userUsecase struct {
	userRepo   *repository.UserRepository
	jwtService *auth.JWTService
}

func NewUserUsecase(userRepo *repository.UserRepository, jwtService *auth.JWTService) UserUsecaseInterface {
	return &userUsecase{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (u *userUsecase) CreateUser(r *request.CreateUserRequest) (*response.AuthResponse, error) {
	user := &domain.User{
		FullName: r.FullName,
		Email:    r.Email,
		Password: pkg.HashPassword(r.Password),
	}

	err := u.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	token, err := u.jwtService.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	authResponse := &response.AuthResponse{
		Token:     token,
		TokenType: "Bearer",
		ExpiresIn: 86400, // 24 hours in seconds
		User: response.UserResponse{
			ID:          user.ID,
			FullName:    user.FullName,
			Email:       user.Email,
			IsConfirmed: user.IsConfirmed,
		},
	}

	return authResponse, nil
}

func (u *userUsecase) LoginUser(r *request.LoginUserRequest) (*response.AuthResponse, error) {
	user, err := u.userRepo.LoginUser(r.Email, r.Password)
	if err != nil {
		return nil, err
	}

	token, err := u.jwtService.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	authResponse := &response.AuthResponse{
		Token:     token,
		TokenType: "Bearer",
		ExpiresIn: 86400, // 24 hours in seconds
		User: response.UserResponse{
			ID:          user.ID,
			FullName:    user.FullName,
			Email:       user.Email,
			IsConfirmed: user.IsConfirmed,
		},
	}

	return authResponse, nil
}

func (u *userUsecase) GetUserByID(userID uint) (*domain.User, error) {
	return u.userRepo.GetUserByID(userID)
}
