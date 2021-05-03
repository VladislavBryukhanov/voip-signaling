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
	router.HandleFunc("/active-connection", connectionmanager.GetActiveConnections).Methods("GET")
	router.HandleFunc("/save-connection", connectionmanager.UpsertConnection).Methods("PUT")

	http.Handle("/", router)
	utils.ErrorHandler(http.ListenAndServe(":3131", nil))
}