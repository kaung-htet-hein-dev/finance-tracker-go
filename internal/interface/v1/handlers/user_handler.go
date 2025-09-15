package handlers

import (
	"errors"
	"fmt"
	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1/request"
	"kaung-htet-hein-dev/finance-tracker-go/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	CreateUser(c echo.Context) error
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

func (h *userHandler) CreateUser(c echo.Context) error {
	u := new(request.CreateUserRequest)
	if err := c.Bind(u); err != nil {
		return c.JSON(400, echo.Map{"error": "Invalid request"})
	}

	if err := c.Validate(u); err != nil {
		return c.JSON(400, echo.Map{"errors": FormatValidationError(err)})
	}

	return c.JSON(201, echo.Map{"email": u.Email})
}

type ValidationErrors struct {
	Field   string
	Message string
}

func FormatValidationError(err error) []ValidationErrors {
	var ve validator.ValidationErrors
	if ok := errors.As(err, &ve); !ok {
		return []ValidationErrors{{Message: err.Error()}}
	}

	errArr := make([]ValidationErrors, 0, len(ve))
	for _, fe := range ve {
		errArr = append(errArr, ValidationErrors{
			Field:   fe.Field(),
			Message: GetErrorMessage(fe),
		})
	}
	return errArr
}

func GetErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Value must be at least %s characters long", fe.Param())
	case "max":
		return fmt.Sprintf("Value must be at most %s characters long", fe.Param())
	}
	return fe.Error()
}
