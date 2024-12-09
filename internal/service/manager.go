package service

import (
	"weather-service/internal/config"
	"weather-service/internal/repo"
)

type service struct {
	repo repo.RepoI
	conf *config.Config
}

type ServiceI interface {
	Weather
}

func New(r repo.RepoI, conf *config.Config) ServiceI {
	return &service{r, conf}
}
