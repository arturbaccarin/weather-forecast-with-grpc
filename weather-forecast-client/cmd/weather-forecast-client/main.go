package main

import (
	"context"
	"fmt"
	"log"

	"github.com/arturbaccarin/weather-forecast-with-grpc/weather-forecast-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var webPort = "9090"

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", webPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}

	client := pb.NewWeatherClient(conn)

	req := &pb.WeatherRequest{
		City: "Brasilia",
	}

	res, err := client.GetWeather(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(res.LocatTime)
}
