package service

import (
	"encoding/json"
	"fmt"
	"weather-service/models"

	"github.com/go-resty/resty/v2"
)

type Weather interface {
	FetchWeather(city string) (*models.Weather, error)
	GetWeather(city string) (*models.Weather, error)
}

func (s *service) GetWeather(city string) (*models.Weather, error) {
	weather, err := s.repo.Get(city)
	if err != nil {
		return nil, err
	}
	return weather, nil
}

func (s *service) FetchWeather(city string) (*models.Weather, error) {
	client := resty.New()

	response, err := client.R().
		SetQueryParams(map[string]string{
			"q":     city,
			"appid": s.conf.WeatherAPI,
			"units": "metric",
		}).Get(s.conf.WeatherBase)

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

	weatherReport := &models.Weather{
		Temperature: data.Main.Temp,
		Humidity:    data.Main.Humidity,
		Description: data.Weather[0].Description,
		City:        city,
	}

	err = s.repo.Add(city, weatherReport)
	if err != nil {
		return nil, err
	}

	return weatherReport, nil
}
