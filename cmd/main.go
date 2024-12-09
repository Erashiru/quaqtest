package main

import (
	"log"
	"net/http"
	"weather-service/internal/config"
	"weather-service/internal/handlers"
	"weather-service/internal/repo"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()

	repo.ConnectMongoDB(config.MongoURI, config.DataBase)

	r := mux.NewRouter()
	r.HandleFunc("/weather", handlers.GetWeather).Methods("GET")
	r.HandleFunc("/weather", handlers.UpdateWeather).Methods("PUT")

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
