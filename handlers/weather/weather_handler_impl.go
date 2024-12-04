package weather

import (
	"fmt"
	"go-agreenery/constants"
	"go-agreenery/dto/base"
	"go-agreenery/dto/response"
	"go-agreenery/helpers"
	"go-agreenery/services/weather"

	"github.com/labstack/echo/v4"
)

type weatherHandler struct {
	service weather.WeatherService
}

func NewWeatherHandler(service weather.WeatherService) *weatherHandler {
	return &weatherHandler{
		service: service,
	}
}

func (handler weatherHandler) GetCurrentWeather(c echo.Context) error {
	lat := c.QueryParam("lat")
	lon := c.QueryParam("lon")
	region := c.QueryParam("region")
	if region != "" {
		lt, lng, err := helpers.GetCoordinates(region)
		if err != nil {
			return base.ErrorResponse(c, err)
		}

		lat = fmt.Sprintf("%f", lt)
		lon = fmt.Sprintf("%f", lng)
	}

	currWeather, err := handler.service.GetCurrentWeather(lat, lon)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetCurrentWeatherSuccess, response.WeatherResponse{}.FromEntity(currWeather))
}

func (handler weatherHandler) GetTodayForecast(c echo.Context) error {
	lat := c.QueryParam("lat")
	lon := c.QueryParam("lon")
	region := c.QueryParam("region")
	if region != "" {
		lt, lng, err := helpers.GetCoordinates(region)
		if err != nil {
			return base.ErrorResponse(c, err)
		}

		lat = fmt.Sprintf("%f", lt)
		lon = fmt.Sprintf("%f", lng)
	}

	forecasts, err := handler.service.GetTodayForecast(lat, lon)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetTodayWeatherSuccess, response.ListWeatherResponse{}.FromListEntity(forecasts))
}

func (handler weatherHandler) GetDailyForecast(c echo.Context) error {
	lat := c.QueryParam("lat")
	lon := c.QueryParam("lon")
	region := c.QueryParam("region")
	if region != "" {
		lt, lng, err := helpers.GetCoordinates(region)
		if err != nil {
			return base.ErrorResponse(c, err)
		}

		lat = fmt.Sprintf("%f", lt)
		lon = fmt.Sprintf("%f", lng)
	}

	forecasts, err := handler.service.GetDailyForecast(lat, lon)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetDailyForecastSuccess, response.ListWeatherResponse{}.FromListEntity(forecasts))
}
