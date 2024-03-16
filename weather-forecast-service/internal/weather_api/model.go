package weatherapi

type WeatherAPIResponse struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Latitude       float64 `json:"lat"`
		Longitude      float64 `json:"lon"`
		TimeZoneID     string  `json:"tz_id"`
		LocalTimeEpoch int64   `json:"localtime_epoch"`
		LocalTime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int64   `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TemperatureC     float64 `json:"temp_c"`
		TemperatureF     float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindSpeedMph    float64 `json:"wind_mph"`
		WindSpeedKph    float64 `json:"wind_kph"`
		WindDegree      int     `json:"wind_degree"`
		WindDirection   string  `json:"wind_dir"`
		PressureMB      float64 `json:"pressure_mb"`
		PressureIn      float64 `json:"pressure_in"`
		PrecipMM        float64 `json:"precip_mm"`
		PrecipIn        float64 `json:"precip_in"`
		Humidity        int     `json:"humidity"`
		Cloud           int     `json:"cloud"`
		FeelsLikeC      float64 `json:"feelslike_c"`
		FeelsLikeF      float64 `json:"feelslike_f"`
		VisibilityKM    float64 `json:"vis_km"`
		VisibilityMiles float64 `json:"vis_miles"`
		UV              float64 `json:"uv"`
		GustMph         float64 `json:"gust_mph"`
		GustKph         float64 `json:"gust_kph"`
	} `json:"current"`
}
