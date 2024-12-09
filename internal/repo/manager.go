package repo

import (
	"weather-service/internal/repo/mongodb"
)

type RepoI interface {
	mongodb.Base
}

func New(uri, dbName string) (RepoI, error) {
	return mongodb.ConnectMongoDB(uri, dbName)
}
