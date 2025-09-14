package server

import (
	"kaung-htet-hein-dev/finance-tracker-go/internal/infrastructure/db"
	v1 "kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1"
	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1/middleware"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	middleware.RegisterBasicMiddleware(e)

	db := db.ConnectDB()

	v1.RegisterUserRoutes(e, db)

	port := ":" + os.Getenv("APP_PORT")

	log.Fatal(e.Start(port))

}
