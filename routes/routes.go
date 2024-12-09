package routes

import (
	authHandler "go-agreenery/handlers/auth"
	categoryHandler "go-agreenery/handlers/category"
	commentHandler "go-agreenery/handlers/comment"
	enrollmentHandler "go-agreenery/handlers/enrollment"
	plantHandler "go-agreenery/handlers/plant"
	postHandler "go-agreenery/handlers/post"
	regionHandler "go-agreenery/handlers/region"
	stepHandler "go-agreenery/handlers/step"
	weatherHandler "go-agreenery/handlers/weather"
	"go-agreenery/middlewares"
	authRepo "go-agreenery/repositories/auth"
	categoryRepo "go-agreenery/repositories/category"
	commentRepo "go-agreenery/repositories/comment"
	enrollmentRepo "go-agreenery/repositories/enrollment"
	plantRepo "go-agreenery/repositories/plant"
	postRepo "go-agreenery/repositories/post"
	stepRepo "go-agreenery/repositories/step"
	authService "go-agreenery/services/auth"
	categoryService "go-agreenery/services/category"
	commentService "go-agreenery/services/comment"
	enrollmentService "go-agreenery/services/enrollment"
	plantService "go-agreenery/services/plant"
	postService "go-agreenery/services/post"
	regionService "go-agreenery/services/region"
	stepService "go-agreenery/services/step"
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

	initPlantRoute(e, db, jwtMiddlewareConfig)

	initStepRoute(e, db, jwtMiddlewareConfig)

	initEnrollmentRoute(e, db, jwtMiddlewareConfig)

	initPostRoute(e, db, jwtMiddlewareConfig)

	initCommentRoute(e, db, jwtMiddlewareConfig)
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
	category.POST("", handler.CreateCategory, middlewares.AdminOnly())
	category.GET("", handler.GetCategories)
	category.GET("/:id", handler.GetCategory)
	category.PUT("/:id", handler.UpdateCategory, middlewares.AdminOnly())
	category.DELETE("/:id", handler.DeleteCategory, middlewares.AdminOnly())
}

func initPlantRoute(e *echo.Echo, db *gorm.DB, jwtMiddlewareConfig echojwt.Config) {
	repository := plantRepo.NewPlantRepository(db)
	service := plantService.NewPlantService(repository)
	handler := plantHandler.NewPlantHandler(service)

	plant := e.Group("/api/v1/plants", echojwt.WithConfig(jwtMiddlewareConfig))
	plant.POST("", handler.CreatePlant, middlewares.AdminOnly())
	plant.GET("", handler.GetPlants)
	plant.GET("/:id", handler.GetPlant)
	plant.PUT("/:id", handler.UpdatePlant, middlewares.AdminOnly())
	plant.DELETE("/:id", handler.DeletePlant, middlewares.AdminOnly())
}

func initStepRoute(e *echo.Echo, db *gorm.DB, jwtMiddlewareConfig echojwt.Config) {
	repository := stepRepo.NewStepRepository(db)
	service := stepService.NewStepService(repository)
	handler := stepHandler.NewStepHandler(service)

	plant := e.Group("/api/v1/steps", echojwt.WithConfig(jwtMiddlewareConfig), middlewares.AdminOnly())
	plant.POST("", handler.CreateStep)
	plant.PUT("/:id", handler.UpdateStep)
	plant.DELETE("/:id", handler.DeleteStep)
}

func initEnrollmentRoute(e *echo.Echo, db *gorm.DB, jwtMiddlewareConfig echojwt.Config) {
	repository := enrollmentRepo.NewEnrollmentRepository(db)
	service := enrollmentService.NewEnrollmentService(repository)
	handler := enrollmentHandler.NewEnrollmentHandler(service)

	enrollment := e.Group("/api/v1/enrollments", echojwt.WithConfig(jwtMiddlewareConfig))
	enrollment.POST("", handler.CreateEnrollment)
	enrollment.GET("", handler.GetEnrollments)
	enrollment.GET("/:enrollmentID", handler.GetEnrollment)
	enrollment.POST("/steps/:stepID/complete", handler.MarkStepAsComplete)
	enrollment.POST("/:enrollmentID/done", handler.SetEnrollmentStatusAsDone)
	enrollment.DELETE("/:enrollmentID", handler.DeleteEnrollment)
}

func initPostRoute(e *echo.Echo, db *gorm.DB, jwtMiddlewareConfig echojwt.Config) {
	repository := postRepo.NewPostRepository(db)
	service := postService.NewPostService(repository)
	handler := postHandler.NewPostHandler(service)

	post := e.Group("/api/v1/posts", echojwt.WithConfig(jwtMiddlewareConfig))
	post.POST("", handler.CreatePost)
	post.GET("", handler.GetPosts)
	post.GET("/:id", handler.GetPost)
	post.PUT("/:id", handler.UpdatePost)
	post.DELETE("/:id", handler.DeletePost)
	post.POST("/:id/like", handler.LikePost)
}

func initCommentRoute(e *echo.Echo, db *gorm.DB, jwtMiddlewareConfig echojwt.Config) {
	repository := commentRepo.NewCommentRepository(db)
	service := commentService.NewCommentService(repository)
	handler := commentHandler.NewCommentHandler(service)

	comment := e.Group("/api/v1/posts/:postID/comments", echojwt.WithConfig(jwtMiddlewareConfig))
	comment.POST("", handler.CreateComment)
	comment.GET("", handler.GetComments)
	comment.PUT("/:id", handler.UpdateComment)
	comment.DELETE("/:id", handler.DeleteComment)
}
