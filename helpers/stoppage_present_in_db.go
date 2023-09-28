package helpers

import (
	"city-bus/configs"
	"city-bus/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var StoppageCollection *mongo.Collection = configs.GetCollection(configs.DB, "stoppage")

func GetStoppage(startingPointName string) (models.Stoppage, error) {
	filter := bson.D{{Key: "stoppagename", Value: startingPointName}}
	var result models.Stoppage
	err := StoppageCollection.FindOne(context.TODO(), filter).Decode(&result)
	return result, err
}
