package v1

import (
	"kaung-htet-hein-dev/finance-tracker-go/internal/infrastructure/repository"
	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1/handlers"
	"kaung-htet-hein-dev/finance-tracker-go/internal/usecase"
	"kaung-htet-hein-dev/finance-tracker-go/pkg"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterUserRoutes(e *echo.Echo, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handlers.NewUserHandler(userUsecase)

	userGroup := e.Group("/api/v1/users")
	userGroup.POST("/register", pkg.BindAndValidate(userHandler.CreateUser))
}
