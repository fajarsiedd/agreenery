package entities

type Weather struct {
	Main struct {
		Temp      float64
		FeelsLike float64
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int
		SeaLevel  int
		GrndLevel int
		Humidity  int
		TempKf    float64
	}
	Weather []struct {
		ID          int
		Main        string
		Description string
		Icon        string
	}
	Wind struct {
		Speed float64
		Deg   int
		Gust  float64
	}
	DtTxt string `json:"dt_txt"`
}
