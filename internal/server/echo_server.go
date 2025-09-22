package server

import (
	"kaung-htet-hein-dev/finance-tracker-go/internal/config"
	"kaung-htet-hein-dev/finance-tracker-go/internal/infrastructure/auth"
	"kaung-htet-hein-dev/finance-tracker-go/internal/infrastructure/db"
	v1 "kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1"
	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1/middleware"
	"kaung-htet-hein-dev/finance-tracker-go/pkg"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func init() {
	cfg = config.LoadConfig()
}

var cfg *config.Config

func StartServer() {
	e := echo.New()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}
	jwtService := auth.NewJWTService(cfg.JWTSecret)

	middleware.RegisterBasicMiddleware(e)
	middleware.RegisterJWTMiddleware(e, jwtService)

	db := db.ConnectDB()

	v1.RegisterUserRoutes(e, db, jwtService)
	v1.RegisterCategoryRoutes(e, db)
	v1.RegisterTransactionRoutes(e, db)

	port := ":" + os.Getenv("APP_PORT")

	log.Fatal(e.Start(port))
}
