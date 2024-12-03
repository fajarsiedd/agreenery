package routes

import (
	authHandler "go-agreenery/handlers/auth"
	regionHandler "go-agreenery/handlers/region"
	"go-agreenery/middlewares"
	authRepo "go-agreenery/repositories/auth"
	authService "go-agreenery/services/auth"
	regionService "go-agreenery/services/region"
	"os"
	"time"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	e.Use(middleware.CORS())

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

	initAuthRoute(e, db, &jwtConfig, jwtMiddlewareConfig)

	initRegionRoute(e, jwtMiddlewareConfig)
}

func initAuthRoute(e *echo.Echo, db *gorm.DB, jwtConfig *middlewares.JWTConfig, jwtMiddlewareConfig echojwt.Config) {
	jwtRefreshConfig := middlewares.JWTConfig{
		SecretKey:       os.Getenv("JWT_REFRESH_SECRET_KEY"),
		ExpiresDuration: 1 * 24 * 7,
	}
	jwtRefreshMiddlewareConfig := jwtRefreshConfig.NewJWTConfig()

	repository := authRepo.NewAuthRepository(db)
	service := authService.NewAuthService(repository, jwtConfig, &jwtRefreshConfig)
	handler := authHandler.NewAuthHandler(service)

	auth := e.Group("/api/v1/auth")
	auth.POST("/login", handler.Login)
	auth.POST("/register", handler.Register)
	auth.POST("/refresh-token", handler.GetNewTokens, echojwt.WithConfig(jwtRefreshMiddlewareConfig))
	auth.GET("/me", handler.GetProfile, echojwt.WithConfig(jwtMiddlewareConfig))
	auth.PUT("/me", handler.UpdateProfile, echojwt.WithConfig(jwtMiddlewareConfig))
	auth.POST("/me/photo", handler.UploadProfilePhoto, echojwt.WithConfig(jwtMiddlewareConfig))
}

func initRegionRoute(e *echo.Echo, jwtMiddlewareConfig echojwt.Config) {
	service := regionService.NewRegionService()
	handler := regionHandler.NewRegionHandler(service)

	region := e.Group("/api/v1/region", echojwt.WithConfig(jwtMiddlewareConfig))
	region.GET("/provinces", handler.GetProvinces)
	region.GET("/regencies/:code", handler.GetRegencies)
	region.GET("/districts/:code", handler.GetDistricts)
	region.GET("/villages/:code", handler.GetVillages)
}
