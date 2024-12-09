package mongodb

import (
	"context"
	"log"
	"weather-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Base interface {
	Add(city string, weather *models.Weather) error
	Get(city string) (*models.Weather, error)
	ClearCollection(dbName string)
}

func (m *MongoDB) Add(city string, weather *models.Weather) error {
	collection := m.DB.Collection("weather")
	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"city": city},
		bson.M{"$set": weather},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) Get(city string) (*models.Weather, error) {
	collection := m.DB.Collection("weather")
	var weather models.Weather

	err := collection.FindOne(context.Background(), bson.M{"city": city}).Decode(&weather)
	if err == mongo.ErrNoDocuments {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return &weather, nil
}

func (m *MongoDB) ClearCollection(dbName string) {
	collection := m.DB.Collection(dbName)

	_, err := collection.DeleteMany(context.Background(), map[string]interface{}{})
	if err != nil {
		log.Fatalf("Error deleting collection %s: %v", dbName, err)
	} else {
		log.Printf("Collection %s was cleared succesfully", collection)
	}
}
