package main

import (
	"city-bus/routers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routers.UserRoute(router)
	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("Hello World")
}
