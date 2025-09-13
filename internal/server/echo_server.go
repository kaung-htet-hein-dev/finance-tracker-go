package server

import (
	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/http/middleware"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	middleware.RegisterBasicMiddleware(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	port := ":" + os.Getenv("APP_PORT")

	e.Logger.Fatal(e.Start(port))
}
