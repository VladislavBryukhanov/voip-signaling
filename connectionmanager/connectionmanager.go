package connectionmanager

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/VladislavBryukhanov/voip-signaling/model"
)

func GetActiveConnections(w http.ResponseWriter, r *http.Request) {
	res := model.GetActiveConnections()

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func UpsertConnection(w http.ResponseWriter, r *http.Request) {
	var conection model.WebRTCConnection

	err := json.NewDecoder(r.Body).Decode(&conection)

	if err != nil {
		// TODO improve me
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error", err)
		return;
	}

	model.CreateWebRTCConnection(&conection)
}
