package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	WeatherAPIURL string
	WeatherAPIKey string
}

func Load() *config {
	godotenv.Load(".env")

	return &config{
		WeatherAPIURL: os.Getenv("WEATHER_API_URL"),
		WeatherAPIKey: os.Getenv("WEATHER_API_KEY"),
	}
}
