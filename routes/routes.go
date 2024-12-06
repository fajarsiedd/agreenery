package routes

import (
	authHandler "go-agreenery/handlers/auth"
	categoryHandler "go-agreenery/handlers/category"
	regionHandler "go-agreenery/handlers/region"
	weatherHandler "go-agreenery/handlers/weather"
	"go-agreenery/middlewares"
	authRepo "go-agreenery/repositories/auth"
	categoryRepo "go-agreenery/repositories/category"
	authService "go-agreenery/services/auth"
	categoryService "go-agreenery/services/category"
	regionService "go-agreenery/services/region"
	weatherService "go-agreenery/services/weather"
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

	initWeatherRoute(e, jwtMiddlewareConfig)

	initCategoryRoute(e, db, jwtMiddlewareConfig)
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
	auth.POST("/refresh", handler.GetNewTokens, echojwt.WithConfig(jwtRefreshMiddlewareConfig))
	auth.GET("/me", handler.GetProfile, echojwt.WithConfig(jwtMiddlewareConfig))
	auth.PUT("/me", handler.UpdateProfile, echojwt.WithConfig(jwtMiddlewareConfig))
	auth.POST("/me/photo", handler.UploadProfilePhoto, echojwt.WithConfig(jwtMiddlewareConfig))
}

func initRegionRoute(e *echo.Echo, jwtMiddlewareConfig echojwt.Config) {
	service := regionService.NewRegionService()
	handler := regionHandler.NewRegionHandler(service)

	region := e.Group("/api/v1/regions", echojwt.WithConfig(jwtMiddlewareConfig))
	region.GET("/provinces", handler.GetProvinces)
	region.GET("/regencies/:code", handler.GetRegencies)
	region.GET("/districts/:code", handler.GetDistricts)
	region.GET("/villages/:code", handler.GetVillages)
}

func initWeatherRoute(e *echo.Echo, jwtMiddlewareConfig echojwt.Config) {
	service := weatherService.NewWeatherService()
	handler := weatherHandler.NewWeatherHandler(service)

	weather := e.Group("/api/v1/weathers", echojwt.WithConfig(jwtMiddlewareConfig))
	weather.GET("/current", handler.GetCurrentWeather)
	weather.GET("/today", handler.GetTodayForecast)
	weather.GET("/daily", handler.GetDailyForecast)
}

func initCategoryRoute(e *echo.Echo, db *gorm.DB, jwtMiddlewareConfig echojwt.Config) {
	repository := categoryRepo.NewCategoryRepository(db)
	service := categoryService.NewCategoryService(repository)
	handler := categoryHandler.NewCategoryHandler(service)

	category := e.Group("/api/v1/categories", echojwt.WithConfig(jwtMiddlewareConfig))
	category.POST("", handler.CreateCategory)
	category.GET("", handler.GetCategories)
	category.GET("/:id", handler.GetCategory)
	category.PUT("/:id", handler.UpdateCategory)
	category.DELETE("/:id", handler.DeleteCategory)
}
