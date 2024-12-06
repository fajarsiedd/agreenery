package response

import (
	"go-agreenery/entities"
	"math"
)

type WeatherResponse struct {
	Main        string  `json:"main"`
	Description string  `json:"description"`
	Temp        float64 `json:"temp"`
	TempMin     float64 `json:"temp_min"`
	TempMax     float64 `json:"temp_max"`
	Humidity    int     `json:"humidity"`
	Icon        string  `json:"icon"`
	WindSpeed   float64 `json:"wind_speed"`
	DateTime    string  `json:"date_time,omitempty"`
}

type ListWeatherResponse []WeatherResponse

func (r WeatherResponse) FromEntity(weather entities.Weather) WeatherResponse {
	var mainTxt string
	switch weather.Weather[0].Main {
	case "Thunderstorm":
		mainTxt = "Hujan Badai"
	case "Drizzle":
		mainTxt = "Gerimis"
	case "Rain":
		mainTxt = "Hujan"
	case "Snow":
		mainTxt = "Salju"
	case "Clear":
		mainTxt = "Cerah"
	case "Clouds":
		mainTxt = "Berawan"
	}

	return WeatherResponse{
		Main:        mainTxt,
		Description: weather.Weather[0].Description,
		Temp:        toFixed(weather.Main.Temp-273.15, 1),
		TempMin:     toFixed(weather.Main.TempMin-273.15, 1),
		TempMax:     toFixed(weather.Main.TempMax-273.15, 1),
		Humidity:    weather.Main.Humidity,
		Icon:        "http://openweathermap.org/img/wn/" + weather.Weather[0].Icon + "@2x.png",
		WindSpeed:   weather.Wind.Speed,
		DateTime:    weather.DtTxt,
	}
}

func (lr ListWeatherResponse) FromListEntity(weathers []entities.Weather) ListWeatherResponse {
	var data ListWeatherResponse

	for _, v := range weathers {
		data = append(data, WeatherResponse{}.FromEntity(v))
	}

	return data
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
