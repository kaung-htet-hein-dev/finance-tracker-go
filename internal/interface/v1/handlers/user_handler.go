package handlers

import (
	"fmt"
	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1/request"
	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	CreateUser(c echo.Context, r *request.CreateUserRequest) error
	LoginUser(c echo.Context, r *request.LoginUserRequest) error
	GetCurrentUser(c echo.Context) error
}

type userHandler struct {
	userUsecase usecase.UserUsecaseInterface
}

func NewUserHandler(userUsecase usecase.UserUsecaseInterface) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

func (h *userHandler) LoginUser(c echo.Context, r *request.LoginUserRequest) error {
	authResponse, err := h.userUsecase.LoginUser(r)
	if err != nil {
		return c.JSON(400, echo.Map{"message": err.Error()})
	}
	return c.JSON(200, echo.Map{
		"message": "Login successful",
		"data":    authResponse,
	})
}

func (h *userHandler) CreateUser(c echo.Context, r *request.CreateUserRequest) error {
	// Create User in DB and get auth response
	authResponse, err := h.userUsecase.CreateUser(r)
	if err != nil {
		return c.JSON(400, echo.Map{"message": err.Error()})
	}

	// Send Confirmation Email
	go func() {}()

	return c.JSON(201, echo.Map{
		"message": fmt.Sprintf("User created successfully. Confirmation email sent to %s", r.Email),
		"data":    authResponse,
	})
}

func (h *userHandler) GetCurrentUser(c echo.Context) error {
	userID := c.Get("user_id")
	user, err := h.userUsecase.GetUserByID(userID.(uint))
	if err != nil {
		return c.JSON(400, echo.Map{"message": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "successful",
		"data":    user,
	})
}
