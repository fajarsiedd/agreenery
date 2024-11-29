package routes

import (
	authHandler "go-agreenery/handlers/auth"
	"go-agreenery/middlewares"
	authRepo "go-agreenery/repositories/auth"
	authService "go-agreenery/services/auth"
	"os"
	"time"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	loggerConfig := middlewares.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}

	loggerMiddleware := loggerConfig.Init()

	e.Use(loggerMiddleware)

	e.Use(middleware.Recover())

	rateLimiterConfig := middlewares.RateLimiterConfig{
		Rate:      10,
		Burst:     30,
		ExpiresIn: 3 * time.Minute,
	}

	rateLimiterMiddleware := rateLimiterConfig.Init()

	e.Use(rateLimiterMiddleware)

	customValidator := middlewares.InitValidator()

	e.Validator = customValidator

	e.Pre(middleware.RemoveTrailingSlash())

	jwtConfig := middlewares.JWTConfig{
		SecretKey:       os.Getenv("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}
	jwtMiddlewareConfig := jwtConfig.NewJWTConfig()

	jwtRefreshConfig := middlewares.JWTConfig{
		SecretKey:       os.Getenv("JWT_REFRESH_SECRET_KEY"),
		ExpiresDuration: 1 * 24 * 7,
	}
	jwtRefreshMiddlewareConfig := jwtRefreshConfig.NewJWTConfig()

	repository := authRepo.NewAuthRepository(db)
	service := authService.NewAuthService(repository, &jwtConfig, &jwtRefreshConfig)
	handler := authHandler.NewAuthHandler(service)

	auth := e.Group("/api/v1/auth")
	auth.POST("/login", handler.Login)
	auth.POST("/register", handler.Register)
	auth.POST("/refresh-token", handler.GetNewTokens, echojwt.WithConfig(jwtRefreshMiddlewareConfig))
	auth.GET("/test", func(c echo.Context) error {
		return c.JSON(200, "Hello")
	}, echojwt.WithConfig(jwtMiddlewareConfig))
}
