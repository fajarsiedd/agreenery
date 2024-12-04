package weather

import "go-agreenery/entities"

type WeatherService interface {
	GetCurrentWeather(lat, lon string) (entities.Weather, error)
	GetTodayForecast(lat, lon string) ([]entities.Weather, error)
	GetDailyForecast(lat, lon string) ([]entities.Weather, error)
}
