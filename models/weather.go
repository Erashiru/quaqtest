package models

type Weather struct {
	City        string  `json:"city" bson:"city"`
	Temperature float64 `json:"temperature" bson:"temperature"`
	Humidity    int     `json:"humidity" bson:"humidity"`
	Description string  `json:"description" bson:"description"`
}
