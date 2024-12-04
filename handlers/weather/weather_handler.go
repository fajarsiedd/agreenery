package weather

import "github.com/labstack/echo/v4"

type WeatherHandler interface {
	GetCurrentWeather(c echo.Context) error
	GetTodayForecast(c echo.Context) error
	GetDailyForecast(c echo.Context) error
}
