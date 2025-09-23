package v1

import (
	"kaung-htet-hein-dev/finance-tracker-go/internal/infrastructure/auth"
	"kaung-htet-hein-dev/finance-tracker-go/internal/infrastructure/repository"
	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1/handlers"
	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1/usecase"
	"kaung-htet-hein-dev/finance-tracker-go/pkg"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterUserRoutes(e *echo.Echo, db *gorm.DB, jwtService *auth.JWTService) {
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, jwtService)
	userHandler := handlers.NewUserHandler(userUsecase)

	userGroup := e.Group("/api/v1/users")
	userGroup.POST("/register", pkg.BindAndValidate(userHandler.CreateUser))
	userGroup.POST("/login", pkg.BindAndValidate(userHandler.LoginUser))
	userGroup.GET("/me", userHandler.GetCurrentUser)
}

func RegisterTransactionRoutes(e *echo.Echo, db *gorm.DB) {
	transactionRepo := repository.NewTransactionRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo, categoryRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionUsecase)

	transactionGroup := e.Group("/api/v1/transactions")
	transactionGroup.POST("", pkg.BindAndValidate(transactionHandler.CreateTransaction))
	transactionGroup.GET("", transactionHandler.GetTransactions)
	transactionGroup.GET("/:id", transactionHandler.GetTransactionByID)
	transactionGroup.PUT("/:id", pkg.BindAndValidate(transactionHandler.UpdateTransaction))
	transactionGroup.DELETE("/:id", transactionHandler.DeleteTransaction)
}

func RegisterCategoryRoutes(e *echo.Echo, db *gorm.DB) {
	categoryRepo := repository.NewCategoryRepository(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryUsecase)

	categoryGroup := e.Group("/api/v1/categories")
	categoryGroup.POST("", pkg.BindAndValidate(categoryHandler.CreateCategory))
	categoryGroup.GET("", categoryHandler.GetCategories)
	categoryGroup.GET("/:id", categoryHandler.GetCategoryByID)
	categoryGroup.PUT("/:id", pkg.BindAndValidate(categoryHandler.UpdateCategory))
	categoryGroup.DELETE("/:id", categoryHandler.DeleteCategory)
}
