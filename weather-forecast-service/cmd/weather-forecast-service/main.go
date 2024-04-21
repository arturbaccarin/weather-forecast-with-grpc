package main

import (
	"fmt"
	"log"
	"net"

	"github.com/arturbaccarin/weather-forecast-with-grpc/weather-forecast-service/pb"
	"google.golang.org/grpc"
)

var webPort = "9090"

func main() {
	// config := config.Load()

	// weatherAPI := weatherapi.NewWeatherAPI(config.WeatherAPIURL, config.WeatherAPIKey)

	// weatherAPI.SearchForCity("London")

	log.Printf("running grpc server on port: %s\n", webPort)

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", webPort))
	if err != nil {
		log.Panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterWeatherServer(s, &Server{})

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
