package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	WeatherAPIKey string
}

func Load() *config {
	godotenv.Load(".env")

	return &config{
		WeatherAPIKey: os.Getenv("WEATHER_API_KEY"),
	}
}
