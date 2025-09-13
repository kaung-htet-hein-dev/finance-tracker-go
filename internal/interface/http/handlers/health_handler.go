package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
