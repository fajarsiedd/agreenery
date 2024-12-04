package weather

import (
	"encoding/json"
	"go-agreenery/entities"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type weatherService struct{}

func NewWeatherService() *weatherService {
	return &weatherService{}
}
func (service weatherService) GetCurrentWeather(lat, lon string) (entities.Weather, error) {
	var err error
	var client = &http.Client{}
	var appID string = os.Getenv("OPEN_WEATHER_API_KEY")
	var url string = "https://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lon + "&lang=id&appid=" + appID

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return entities.Weather{}, err
	}

	response, err := client.Do(request)
	if err != nil {
		return entities.Weather{}, err
	}
	defer response.Body.Close()

	var res entities.Weather
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return entities.Weather{}, err
	}

	return res, nil
}

func (service weatherService) GetTodayForecast(lat, lon string) ([]entities.Weather, error) {
	var err error
	var client = &http.Client{}
	var appID string = os.Getenv("OPEN_WEATHER_API_KEY")
	var url string = "https://api.openweathermap.org/data/2.5/forecast?lat=" + lat + "&lon=" + lon + "&lang=id&appid=" + appID

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	type resBody struct {
		List []entities.Weather
	}

	var res resBody
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	var filterredWeathers []entities.Weather
	dt_format := "2006-01-02"
	now := time.Now().Format(dt_format)
	for _, v := range res.List {
		if strings.Contains(v.DtTxt, now) {
			filterredWeathers = append(filterredWeathers, v)
		}
	}

	return filterredWeathers, nil
}

func (service weatherService) GetDailyForecast(lat, lon string) ([]entities.Weather, error) {
	var err error
	var client = &http.Client{}
	var appID string = os.Getenv("OPEN_WEATHER_API_KEY")
	var url string = "https://api.openweathermap.org/data/2.5/forecast?lat=" + lat + "&lon=" + lon + "&lang=id&appid=" + appID

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	type resBody struct {
		List []entities.Weather
	}

	var res resBody
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	filterredWeathers := make([]entities.Weather, 5)
	j := 0
	dateFormat := "2006-01-02"
	date := time.Now().AddDate(0, 0, 1).Format(dateFormat)
	minDiff := 0

	for i := 0; i < len(res.List); i++ {
		if !strings.Contains(res.List[i].DtTxt, time.Now().Format(dateFormat)) {
			if strings.Contains(res.List[i].DtTxt, date) {
				weatherHour, _ := strconv.Atoi(res.List[i].DtTxt[11:13])

				if res.List[i].DtTxt[11:13] == "00" {
					weatherHour = 24
				}

				diff := abs(weatherHour, time.Now().Hour())

				if diff >= 0 && (minDiff == 0 || diff < minDiff) {
					minDiff = diff
					filterredWeathers[j] = res.List[i]
				}
			} else {
				dt, err := time.Parse(dateFormat, date)
				if err != nil {
					return nil, err
				}

				date = dt.AddDate(0, 0, 1).Format(dateFormat)
				minDiff = 0
				i--
				j++
			}
		}
	}

	return filterredWeathers, nil
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
