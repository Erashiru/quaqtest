package handlers

import (
	"weather-service/internal/config"
	"weather-service/internal/service"
)

type handlers struct {
	service service.ServiceI
	config  *config.Config
}

func New(s service.ServiceI, config *config.Config) *handlers {
	return &handlers{s, config}
}
