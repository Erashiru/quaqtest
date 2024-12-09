package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *handlers) Routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/weather", h.GetWeather).Methods("GET")
	r.HandleFunc("/weather", h.UpdateWeather).Methods("PUT")
	return r
}
