package helpers

import (
	"city-bus/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllStoppage(startingPoint, endingPoint string) ([]models.BusPreview, error) {
	var targetStoppage models.Stoppage
	//find the target stoppage starting with starting point
	filter := bson.D{{Key: "stoppagename", Value: startingPoint}}
	err := StoppageCollection.FindOne(context.TODO(), filter).Decode(&targetStoppage)
	if err != nil {
		return targetStoppage.Destination, err
	}
	allDestination := targetStoppage.Destination
	var filteredDestination []models.BusPreview
	for _, destination := range allDestination {
		if destination.EndingPointName == endingPoint {
			filteredDestination = append(filteredDestination, destination)
		}
	}
	return filteredDestination, err
}
