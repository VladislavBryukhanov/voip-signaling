module github.com/VladislavBryukhanov/voip-signaling/model

go 1.16

replace github.com/VladislavBryukhanov/voip-signaling/utils => ../utils

require (
	github.com/VladislavBryukhanov/voip-signaling/utils v0.0.0-00010101000000-000000000000
	github.com/spf13/viper v1.7.1
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.9
)
