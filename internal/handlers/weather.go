package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *handlers) GetWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City parameter is required", http.StatusBadRequest)
		return
	}

	//func get
	weather, err := h.service.GetWeather(city)
	if err != nil {
		http.Error(w, "Couldn't get data from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(weather)
}

func (h *handlers) UpdateWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City parameter is required", http.StatusBadRequest)
		return
	}

	weather, err := h.service.FetchWeather(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(weather)
}
