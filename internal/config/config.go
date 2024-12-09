package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI    string
	DataBase    string
	WeatherAPI  string
	WeatherBase string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", err)
	}

	conf := Config{
		MongoURI:    os.Getenv("MONGO_URI"),
		DataBase:    os.Getenv("DATABASE_NAME"),
		WeatherAPI:  os.Getenv("WEATHER_API_KEY"),
		WeatherBase: "https://api.openweathermap.org/data/2.5/weather",
	}

	return &conf
}
