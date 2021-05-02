module github.com/VladislavBryukhanov/voip-signaling

go 1.16

require (
	github.com/VladislavBryukhanov/voip-signaling/connectionmanager v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)

replace github.com/VladislavBryukhanov/voip-signaling/connectionmanager => ./connectionmanager
