package middleware

import (
	"kaung-htet-hein-dev/finance-tracker-go/internal/infrastructure/auth"
	"net/http"
	"strings"

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

func RegisterJWTMiddleware(e *echo.Echo, jwtService *auth.JWTService) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !shouldValidate(c) {
				return next(c)
			}

			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{"message": "Missing Authorization Header"})
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			if token == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{"message": "Invalid Authorization Header"})
			}

			claims, err := jwtService.ValidateToken(token)

			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{"message": "Invalid Token"})
			}

			c.Set("user_id", claims.UserID)
			c.Set("email", claims.Email)

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
