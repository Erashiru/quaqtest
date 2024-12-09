package service

import (
	"encoding/json"
	"fmt"
	"weather-service/internal/config"
	"weather-service/models"

	"github.com/go-resty/resty/v2"
)

func FetchWeather(city string) (*models.Weather, error) {
	client := resty.New()

	response, err := client.R().
		SetQueryParams(map[string]string{
			"q":     city,
			"appid": config.WeatherAPI,
			"units": "metric",
		}).Get(config.WeatherBase)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to fetch weather data: %s", response.String())
	}

	var data struct {
		Main struct {
			Temp     float64 `json:"temp"`
			Humidity int     `json:"humidity"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
	}
	if err := json.Unmarshal(response.Body(), &data); err != nil {
		return nil, err
	}

	return &models.Weather{
		Temperature: data.Main.Temp,
		Humidity:    data.Main.Humidity,
		Description: data.Weather[0].Description,
		City:        city,
	}, nil
}
