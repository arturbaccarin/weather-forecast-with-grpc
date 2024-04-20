package weatherapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type WeatherAPI struct {
	URL string
	Key string
}

func NewWeatherAPI(url, key string) *WeatherAPI {
	return &WeatherAPI{
		URL: url,
		Key: key,
	}
}

func (w WeatherAPI) SearchForCity(city string) {
	endpoint := "%s?key=%s&q=%s&aqi=no"

	url := fmt.Sprintf(endpoint, w.URL, w.Key, city)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("error: " + err.Error())
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("error: " + err.Error())
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	var weatherResponse WeatherAPIResponse

	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	dto := WeatherOutputDTO{
		Name:               weatherResponse.Location.Name,
		Condition:          weatherResponse.Current.Condition.Text,
		TemperatureCelsius: weatherResponse.Current.TemperatureC,
	}

	fmt.Println(dto)
}
