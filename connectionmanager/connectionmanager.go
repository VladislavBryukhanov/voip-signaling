package connectionmanager

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/VladislavBryukhanov/voip-signaling/model"
	"github.com/gorilla/mux"
)

func httpErrorHandler(err error, w http.ResponseWriter) {
	if err != nil {
		// TODO improve me
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error", err)
		return;
	}
}

func GetActiveConnections(w http.ResponseWriter, r *http.Request) {
	res, err := model.GetActiveConnections()
	httpErrorHandler(err, w)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func UpsertConnection(w http.ResponseWriter, r *http.Request) {
	var conection model.WebRTCConnection

	err := json.NewDecoder(r.Body).Decode(&conection)
	httpErrorHandler(err, w)

	// TODO upsert
	model.CreateWebRTCConnection(&conection)
}

func AttachIceCandidate(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["connection_id"]
	var ice model.IceCandidate

	connId, err := strconv.ParseUint(idParam, 10, 64)
	httpErrorHandler(err, w)

	err = json.NewDecoder(r.Body).Decode(&ice)
	httpErrorHandler(err, w)

	ice.ConnectionID = uint(connId)

	err = model.AttachIceCandidate(&ice)
	httpErrorHandler(err, w)
}

func DisposeConnection(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["connection_id"]

	connectionId, err := strconv.ParseUint(idParam, 10, 64)
	httpErrorHandler(err, w)

	err = model.DeleteConnection(uint(connectionId))
	httpErrorHandler(err, w)
}