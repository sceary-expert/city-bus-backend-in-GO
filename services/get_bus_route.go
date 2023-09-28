package services

import (
	"city-bus/helpers"
	"city-bus/models"
	"city-bus/responses"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func GetBusRoute(rw http.ResponseWriter, r *http.Request) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var requestRoute models.RequestRoute
	if err := json.NewDecoder(r.Body).Decode(&requestRoute); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error : request body should contain every field", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(rw).Encode(response)
		return
	}
	startingPoint := requestRoute.StartingPointName
	endingPoint := requestRoute.EndingPointName
	result, err := helpers.GetAllStoppage(startingPoint, endingPoint)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error : request body should contain every field", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(rw).Encode(response)
		return
	}
	rw.WriteHeader(http.StatusAccepted)
	response := responses.UserResponse{Status: http.StatusBadRequest, Message: "success", Data: map[string]interface{}{"data": result}}
	json.NewEncoder(rw).Encode(response)
	return
}
