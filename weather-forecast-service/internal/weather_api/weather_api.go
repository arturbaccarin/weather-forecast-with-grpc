package weatherapi

import (
	"fmt"
	"io"
	"net/http"
)

type WeatherAPI struct {
	URL string
	Key string
}

func NewWeatherAPI(key string) *WeatherAPI {
	return &WeatherAPI{
		Key: key,
	}
}

func (w WeatherAPI) SearchForCity(city string) {
	endpoint := "%s?key=%s&q=%s&aqi=no"

	url := fmt.Sprintf(endpoint, w.URL, w.Key, city)
	println(url)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()

	responseBody, _ := io.ReadAll(response.Body)

	println(string(responseBody))
}
