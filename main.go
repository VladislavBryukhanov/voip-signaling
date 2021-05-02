package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/VladislavBryukhanov/voip-signaling/connectionmanager"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/active-connection", connectionmanager.GetActiveConnections).Methods("GET")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":3131", nil))
}