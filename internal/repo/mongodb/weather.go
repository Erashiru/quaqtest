package mongodb

import (
	"context"
	"weather-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Weather interface {
	Add(city string, weather *models.Weather) error
	Get(city string) (*models.Weather, error)
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
