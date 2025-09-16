package handlers

import (
	"fmt"
	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1/request"
	"kaung-htet-hein-dev/finance-tracker-go/internal/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	CreateUser(c echo.Context, r *request.CreateUserRequest) error
}

type userHandler struct {
	userUsecase usecase.UserUsecaseInterface
}

func NewUserHandler(userUsecase usecase.UserUsecaseInterface) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

func (h *userHandler) CreateUser(c echo.Context, r *request.CreateUserRequest) error {
	// Create User in DB
	_, err := h.userUsecase.CreateUser(r)

	if err != nil {
		return c.JSON(400, echo.Map{"message": err.Error()})
	}

	// Send Confirmation Email
	go func() {}()

	return c.JSON(201, echo.Map{"email": fmt.Sprintf("Confirmation email has been sent to %s", r.Email)})
}
