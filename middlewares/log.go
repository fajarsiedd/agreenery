package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type LoggerConfig struct {
	Format string
}

func (loggerConfig *LoggerConfig) Init() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: loggerConfig.Format,
	})
}
