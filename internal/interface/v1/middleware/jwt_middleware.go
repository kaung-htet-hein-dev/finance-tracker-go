package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type path struct {
	Method string
	Path   string
}

var apiPrefix = "/api/v1"

var protectedRoutes = []path{
	{Method: http.MethodGet, Path: apiPrefix + "/users/me"},
}

func RegisterJWTMiddleware(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !shouldValidate(c) {
				return next(c)
			}

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{"message": "Missing Authorization Header"})
			}

			// JWT authentication logic here
			return next(c)
		}
	})
}

func shouldValidate(c echo.Context) bool {
	requestMethod := c.Request().Method
	requestPath := c.Request().URL.Path

	for _, route := range protectedRoutes {
		if route.Method == requestMethod && route.Path == requestPath {
			return true
		}
	}
	return false
}
