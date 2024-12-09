package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"weather-service/mocks"
	"weather-service/models"

	"github.com/gorilla/mux"
)

func TestGetWeather(t *testing.T) {
	// Инициализация моков
	mockRepo := mocks.NewMockRepo()
	mockService := mocks.NewMockService(mockRepo)
	h := New(mockService, nil)

	// Добавляем тестовые данные
	mockRepo.Add("Almaty", &models.Weather{
		City:        "Almaty",
		Temperature: 24.8,
		Humidity:    60,
		Description: "clear sky",
	})

	// Создаем запрос
	req, err := http.NewRequest("GET", "/weather?city=Almaty", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/weather", h.GetWeather).Methods("GET")
	router.ServeHTTP(rr, req)

	// Проверяем ответ
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestUpdateWeather(t *testing.T) {
	// Инициализация моков
	mockRepo := mocks.NewMockRepo()
	mockService := mocks.NewMockService(mockRepo)
	h := New(mockService, nil)

	// Создаем запрос
	req, err := http.NewRequest("PUT", "/weather?city=Astana", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/weather", h.UpdateWeather).Methods("PUT")
	router.ServeHTTP(rr, req)

	// Проверяем ответ
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}
