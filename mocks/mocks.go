package mocks

import (
	"errors"
	"weather-service/models"
)

// MockRepo — мок для репозитория
type MockRepo struct {
	Data map[string]*models.Weather
}

func NewMockRepo() *MockRepo {
	return &MockRepo{Data: make(map[string]*models.Weather)}
}

func (m *MockRepo) Get(city string) (*models.Weather, error) {
	weather, exists := m.Data[city]
	if !exists {
		return nil, errors.New("city not found")
	}
	return weather, nil
}

func (m *MockRepo) Add(city string, weather *models.Weather) error {
	m.Data[city] = weather
	return nil
}

// MockService — мок для сервиса
type MockService struct {
	Repo *MockRepo
}

func NewMockService(repo *MockRepo) *MockService {
	return &MockService{Repo: repo}
}

func (m *MockService) GetWeather(city string) (*models.Weather, error) {
	return m.Repo.Get(city)
}

func (m *MockService) FetchWeather(city string) (*models.Weather, error) {
	// Возвращаем фиктивные данные для тестов
	weather := &models.Weather{
		City:        city,
		Temperature: 25.0,
		Humidity:    60,
		Description: "clear sky",
	}
	return weather, m.Repo.Add(city, weather)
}
