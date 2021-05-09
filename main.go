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

func httpHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	});
}

func main() {
	initEnv()
	model.InitDb()
	model.Migrate()

	router := mux.NewRouter().PathPrefix("/connection").Subrouter()
	// TODO how async works

	router.Use(httpHeaderMiddleware);

	router.HandleFunc("", connectionmanager.GetActiveConnections).Methods("GET")
	router.HandleFunc("/{connection_id}", connectionmanager.GetConnection).Methods("GET")
	router.HandleFunc("/{connection_id}", connectionmanager.UpsertConnection).Methods("PUT")
	router.HandleFunc("/{connection_id}", connectionmanager.DisposeConnection).Methods("DELETE")
	router.HandleFunc("/{connection_id}/session-description", connectionmanager.AttachSessionDescription).Methods("PUT")
	router.HandleFunc("/{connection_id}/ice-candidate", connectionmanager.AttachIceCandidate).Methods("POST")

	http.Handle("/", router)
	utils.ErrorHandler(http.ListenAndServe(":3131", nil))
}