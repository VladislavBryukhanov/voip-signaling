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

func getConnectionId(r *http.Request) (uint, error) {
	idParam := mux.Vars(r)["connection_id"]
	connId, err := strconv.ParseUint(idParam, 10, 64)
	return uint(connId), err
}

func GetActiveConnections(w http.ResponseWriter, r *http.Request) {
	res, err := model.GetActiveConnections()
	httpErrorHandler(err, w)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func UpsertConnection(w http.ResponseWriter, r *http.Request) {
	var conection model.WebRTCConnection

	connectionId, err := getConnectionId(r)
	httpErrorHandler(err, w)

	err = json.NewDecoder(r.Body).Decode(&conection)
	httpErrorHandler(err, w)

	conection.ID = connectionId

	err = model.CreateWebRTCConnection(&conection)
	httpErrorHandler(err, w)
}

func AttachIceCandidate(w http.ResponseWriter, r *http.Request) {
	var ice model.IceCandidate

	connectionId, err := getConnectionId(r)
	httpErrorHandler(err, w)

	err = json.NewDecoder(r.Body).Decode(&ice)
	httpErrorHandler(err, w)

	ice.ConnectionID = connectionId

	err = model.AttachIceCandidate(&ice)
	httpErrorHandler(err, w)
}

func AttachSessionDescription(w http.ResponseWriter, r *http.Request) {
	var description model.RTCSessionDescription

	connectionId, err := getConnectionId(r)
	httpErrorHandler(err, w)

	err = json.NewDecoder(r.Body).Decode(&description)
	httpErrorHandler(err, w)

	description.ConnectionID = connectionId

	err = model.AttachSessionDescription(&description)
	httpErrorHandler(err, w)
}

func DisposeConnection(w http.ResponseWriter, r *http.Request) {
	connectionId, err := getConnectionId(r)
	httpErrorHandler(err, w)

	err = model.DeleteConnection(connectionId)
	httpErrorHandler(err, w)
}