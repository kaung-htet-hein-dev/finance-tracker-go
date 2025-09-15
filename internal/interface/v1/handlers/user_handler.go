package handlers

import (
	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1/request"
	"kaung-htet-hein-dev/finance-tracker-go/internal/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	CreateUser(c echo.Context, r *request.CreateUserRequest) error
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

func (h *userHandler) CreateUser(c echo.Context, r *request.CreateUserRequest) error {

	return c.JSON(201, echo.Map{"email": r.Email})
}
