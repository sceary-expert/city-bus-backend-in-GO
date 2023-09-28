package helpers

import (
	"city-bus/models"
	"context"
	"fmt"
	"time"
)

func CreateNewStoppage(startingPointName string) (models.Stoppage, error) {
	fmt.Println("inside create new stoppage")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var newStoppage models.Stoppage
	newStoppage.StoppageName = startingPointName
	fmt.Println("before creating new stoppage")
	_, err := StoppageCollection.InsertOne(ctx, newStoppage)
	fmt.Println("after creating new stoppage")
	fmt.Println("newStoppage ", newStoppage)
	if err != nil {
		return newStoppage, err
	}
	return newStoppage, err
}
