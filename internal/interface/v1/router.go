package v1

import (
	"kaung-htet-hein-dev/finance-tracker-go/internal/config"
	"kaung-htet-hein-dev/finance-tracker-go/internal/infrastructure/auth"
	"kaung-htet-hein-dev/finance-tracker-go/internal/infrastructure/repository"
	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1/handlers"
	"kaung-htet-hein-dev/finance-tracker-go/internal/usecase"
	"kaung-htet-hein-dev/finance-tracker-go/pkg"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func init() {
	cfg = config.LoadConfig()
}

var cfg *config.Config

func RegisterUserRoutes(e *echo.Echo, db *gorm.DB) {
	jwtService := auth.NewJWTService(cfg.JWTSecret)
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, jwtService)
	userHandler := handlers.NewUserHandler(userUsecase)

	userGroup := e.Group("/api/v1/users")
	userGroup.POST("/register", pkg.BindAndValidate(userHandler.CreateUser))
	userGroup.POST("/login", pkg.BindAndValidate(userHandler.LoginUser))
	userGroup.GET("/me", userHandler.GetCurrentUser)
}

func RegisterTransactionRoutes(e *echo.Echo, db *gorm.DB) {}
