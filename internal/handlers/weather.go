package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"weather-service/internal/repo"
	"weather-service/internal/service"
	"weather-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City parameter is required", http.StatusBadRequest)
		return
	}

	collection := repo.DB.Collection("weather")
	var weather models.Weather

	err := collection.FindOne(context.Background(), bson.M{"city": city}).Decode(&weather)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "Weather data not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(weather)
}

func UpdateWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City parameter is required", http.StatusBadRequest)
		return
	}

	weather, err := service.FetchWeather(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	collection := repo.DB.Collection("weather")
	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{"city": city},
		bson.M{"$set": weather},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(weather)
}
