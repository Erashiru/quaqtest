package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	DB *mongo.Database
}

func ConnectMongoDB(uri, dbName string) (*MongoDB, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	DB := client.Database(dbName)
	return &MongoDB{DB: DB}, nil
}
