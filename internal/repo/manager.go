package repo

import (
	"weather-service/internal/repo/mongodb"
)

type RepoI interface {
	mongodb.Weather
}

func New(uri, dbName string) (RepoI, error) {
	return mongodb.ConnectMongoDB(uri, dbName)
}
