module github.com/VladislavBryukhanov/voip-signaling

go 1.16

require (
	github.com/VladislavBryukhanov/voip-signaling/connectionmanager v0.0.0-00010101000000-000000000000
	github.com/VladislavBryukhanov/voip-signaling/model v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.9
)

replace github.com/VladislavBryukhanov/voip-signaling/connectionmanager => ./connectionmanager

replace github.com/VladislavBryukhanov/voip-signaling/model => ./model
