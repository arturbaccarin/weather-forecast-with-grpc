package main

import (
	"github.com/arturbaccarin/weather-forecast-with-grpc/weather-forecast-service/config"
	weatherapi "github.com/arturbaccarin/weather-forecast-with-grpc/weather-forecast-service/internal/weather_api"
)

func main() {
	println("start project")

	config := config.Load()

	weatherAPI := weatherapi.NewWeatherAPI(config.WeatherAPIKey)

	weatherAPI.SearchForCity("London")

}
