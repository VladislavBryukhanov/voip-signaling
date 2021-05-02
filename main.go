package main

import (
	"log"
	"net/http"
	"github.com/VladislavBryukhanov/voip-signaling/connectionmanager"
	"github.com/VladislavBryukhanov/voip-signaling/model"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DSN = "host=localhost user=SECRET password=SECRET dbname=voip_signaling port=5432"

func initDb() {
	db, err := gorm.Open(postgres.Open(DSN))

	if err != nil {
		log.Fatal(err)
	}
	model.SetDatabase(db, true)
}

func main() {
	initDb();
	
	router := mux.NewRouter()
	router.HandleFunc("/mock-connection", connectionmanager.GetMockConnections).Methods("GET")
	router.HandleFunc("/active-connection", connectionmanager.GetActiveConnections).Methods("GET")
	router.HandleFunc("/save-connection", connectionmanager.UpsertConnection).Methods("PUT")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":3131", nil))
}