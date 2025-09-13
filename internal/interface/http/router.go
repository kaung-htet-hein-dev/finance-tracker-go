package http

import (
	echo "github.com/labstack/echo/v4"

	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/http/handlers"
	imw "kaung-htet-hein-dev/finance-tracker-go/internal/interface/http/middleware"
)

// Register wires middleware and routes onto the provided Echo instance.
func Register(e *echo.Echo) {
	imw.RegisterBasicMiddleware(e)

	// Routes
	e.GET("/health", handlers.Health)
}
