package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func RegisterBasicMiddleware(e *echo.Echo) {
	e.HideBanner = true
	e.Logger.SetLevel(log.INFO)

	// Pre-middlewares
	e.Pre(middleware.RemoveTrailingSlash())

	// Core middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// Includes request id, remote ip, method, uri, status, latency, sizes and UA
		Format: "${method} ${uri} | status=${status} latency=${latency_human} in=${bytes_in}B out=${bytes_out}B | ua=\"${user_agent}\"\n",
	}))
}
