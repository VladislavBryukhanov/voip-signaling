package main

import (
	"net/http"
	"github.com/VladislavBryukhanov/voip-signaling/connectionmanager"
	"github.com/VladislavBryukhanov/voip-signaling/model"
	"github.com/VladislavBryukhanov/voip-signaling/utils"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func initEnv() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig();
	utils.ErrorHandler(err)
}

func main() {
	initEnv()
	model.InitDb()
	model.Migrate()

	router := mux.NewRouter()
	// TODO how async works

	router.HandleFunc("/connection", connectionmanager.GetActiveConnections).Methods("GET")
	router.HandleFunc("/connection/{connection_id}", connectionmanager.UpsertConnection).Methods("PUT")
	router.HandleFunc("/connection/{connection_id}", connectionmanager.DisposeConnection).Methods("DELETE")
	router.HandleFunc("/connection/{connection_id}/session-description", connectionmanager.AttachSessionDescription).Methods("PUT")
	router.HandleFunc("/connection/{connection_id}/ice-candidate", connectionmanager.AttachIceCandidate).Methods("POST")

	http.Handle("/", router)
	utils.ErrorHandler(http.ListenAndServe(":3131", nil))
}