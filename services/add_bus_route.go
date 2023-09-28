package services

import (
	"city-bus/helpers"
	"city-bus/models"
	"city-bus/responses"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func AddBusRoute(rw http.ResponseWriter, r *http.Request) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var newRouteRequstBody models.NewRouteRequstBody
	if err := json.NewDecoder(r.Body).Decode(&newRouteRequstBody); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error : request body should contain every field", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(rw).Encode(response)
		return
	}
	startingPointName := newRouteRequstBody.StartingPointName

	var stoppage models.Stoppage

	//find starting point

	stoppage, err := helpers.GetStoppage(startingPointName)
	fmt.Println("stoppage result from get stoppage ", stoppage)
	if err != nil {
		//if no record in database create a new starting point
		stoppage, err = helpers.CreateNewStoppage(startingPointName)
		fmt.Println("new stoppage created ", stoppage)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error : request body should contain every field", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

	}
	fmt.Println("stoppage result after creating a new stoppage ", stoppage)
	var newDestination models.BusPreview
	newDestination.BusName = newRouteRequstBody.BusName
	newDestination.Duration = newRouteRequstBody.Duration
	newDestination.EndingPointName = newRouteRequstBody.EndingPointName
	newDestination.TicketPrice = newRouteRequstBody.TicketPrice
	destinationList := stoppage.Destination
	destinationList = append(destinationList, newDestination)
	stoppage.Destination = destinationList
	err = helpers.UpdateStoppage(stoppage)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error : not able to add new destination", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(rw).Encode(response)
		return
	}
	rw.WriteHeader(http.StatusCreated)
	response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error : request body should contain every field", Data: map[string]interface{}{"data": newDestination}}
	json.NewEncoder(rw).Encode(response)
	return

}
