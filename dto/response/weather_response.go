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

func (weather WeatherResponse) FromEntity(weatherEntity entities.Weather) WeatherResponse {
	var mainTxt string
	switch weatherEntity.Weather[0].Main {
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
		Description: weatherEntity.Weather[0].Description,
		Temp:        toFixed(weatherEntity.Main.Temp-273.15, 1),
		TempMin:     toFixed(weatherEntity.Main.TempMin-273.15, 1),
		TempMax:     toFixed(weatherEntity.Main.TempMax-273.15, 1),
		Humidity:    weatherEntity.Main.Humidity,
		Icon:        "http://openweathermap.org/img/wn/" + weatherEntity.Weather[0].Icon + "@2x.png",
		WindSpeed:   weatherEntity.Wind.Speed,
		DateTime:    weatherEntity.DtTxt,
	}
}

func (listWeather ListWeatherResponse) FromListEntity(weatherEntities []entities.Weather) ListWeatherResponse {
	var weathers ListWeatherResponse

	for _, v := range weatherEntities {
		weathers = append(weathers, WeatherResponse{}.FromEntity(v))
	}

	return weathers
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
