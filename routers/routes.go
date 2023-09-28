package routers

import (
	"city-bus/services"

	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/add-bus-route", services.AddBusRoute).Methods("POST")
	// fmt.Println("routes 9")
	// router.HandleFunc("/set", commands.Set()).Methods("POST")
	// router.HandleFunc("/get", commands.Get()).Methods("POST")

}
