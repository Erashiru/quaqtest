package main

import (
	"log"
	"net/http"
	"weather-service/internal/config"
	"weather-service/internal/handlers"
	"weather-service/internal/repo"
	"weather-service/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	config := config.LoadConfig()

	repo, err := repo.New(config.MongoURI, config.DataBase)
	if err != nil {
		log.Fatalf("Error with database connection: %v", err)
	}

	s := service.New(repo, config)
	h := handlers.New(s, config)

	r := mux.NewRouter()
	r.HandleFunc("/weather", h.GetWeather).Methods("GET")
	r.HandleFunc("/weather", h.UpdateWeather).Methods("PUT")

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
