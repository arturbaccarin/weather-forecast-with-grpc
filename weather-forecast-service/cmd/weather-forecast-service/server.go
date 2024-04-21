package main

import (
	"context"

	"github.com/arturbaccarin/weather-forecast-with-grpc/weather-forecast-service/pb"
)

type Server struct {
	pb.UnimplementedWeatherServer
}

func (s *Server) GetWeather(ctx context.Context, in *pb.WeatherRequest) (*pb.WeatherResponse, error) {
	localTime := "2024 04 21"

	return &pb.WeatherResponse{LocatTime: localTime}, nil
}
