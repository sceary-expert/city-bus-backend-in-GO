package helpers

import (
	"city-bus/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateStoppage(stoppage models.Stoppage) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//TODO : instead of searching by name try to search with id
	_, err := StoppageCollection.UpdateOne(ctx, bson.M{"stoppagename": stoppage.StoppageName}, bson.M{"$set": stoppage})
	return err
}
