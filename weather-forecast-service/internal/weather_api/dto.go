package weatherapi

type WeatherOutputDTO struct {
	Name               string  `json:"name"`
	Condition          string  `json:"condition"`
	TemperatureCelsius float64 `json:"temp_c"`
}
