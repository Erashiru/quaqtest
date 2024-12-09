package main

import (
	"log"
	"net/http"
	"weather-service/internal/config"
	"weather-service/internal/handlers"
	"weather-service/internal/repo"
	"weather-service/internal/service"
)

func main() {
	config := config.LoadConfig()

	repo, err := repo.New(config.MongoURI, config.DataBase)
	if err != nil {
		log.Fatalf("Error with database connection: %v", err)
	}
	s := service.New(repo, config)
	h := handlers.New(s, config)

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", h.Routes()))
}
