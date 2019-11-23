package main

import (
	"encoding/json"
	"fmt"
	"go-api/data"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func printAllEvents() {

	listOfEvents := data.GetAllEvents()

	for _, value := range listOfEvents {
		str, _ := json.Marshal(value)
		fmt.Println(string(str))
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	listOfEvents := data.GetAllEvents()
	json.NewEncoder(w).Encode(listOfEvents)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/events", get).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":80", r))
}
