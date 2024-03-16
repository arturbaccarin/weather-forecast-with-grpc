package main

import (
	"github.com/arturbaccarin/weather-forecast-with-grpc/weather-forecast-service/config"
	weatherapi "github.com/arturbaccarin/weather-forecast-with-grpc/weather-forecast-service/internal/weather_api"
)

func main() {
	config := config.Load()

	weatherAPI := weatherapi.NewWeatherAPI(config.WeatherAPIURL, config.WeatherAPIKey)

	weatherAPI.SearchForCity("London")

}
