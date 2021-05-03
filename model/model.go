package model

import (
	"time"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/VladislavBryukhanov/voip-signaling/utils"
)

type RTCSessionDescriptionInit struct {
	Sdp string `json:"sdp"`
	Type string `json:"type"` 
}

type IceCandidate struct {
	Candidate string `json:"candidate"`
	SdpMLineIndex int `json:"sdpMLineIndex"`
	SdpMid int `json:"sdpMid"`
}

type WebRTCConnection struct {
	ID uint `gorm:"primaryKey"`
	CreatedAt time.Time

	InitiatorId int
	ExpirationDate int 
	Offer RTCSessionDescriptionInit `gorm:"embedded"`
	Answer RTCSessionDescriptionInit `gorm:"embedded"`
	// Candidates []IceCandidate `gorm:"embedded" json:"cadidates"`
}

var DB *gorm.DB

func InitDb() {
	dsn := viper.Get("DSN").(string)

	db, err := gorm.Open(postgres.Open(dsn))
	utils.ErrorHandler(err)
	DB = db
}

func Migrate() {
	DB.AutoMigrate(&WebRTCConnection{})
}

func CreateWebRTCConnection(con *WebRTCConnection) {
	res := DB.Create(con)
	utils.ErrorHandler(res.Error)
}

func GetActiveConnections() []WebRTCConnection {
	var peerCons []WebRTCConnection

	res := DB.Find(&peerCons)
	utils.ErrorHandler(res.Error)

	return peerCons
}