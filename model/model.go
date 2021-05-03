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
	ID uint `gorm:"primaryKey; uniqueIndex"`
	ConnectionID uint

	Candidate string `json:"candidate"`
	SdpMLineIndex int `json:"sdpMLineIndex"`
	SdpMid int `json:"sdpMid"`
}

// TODO set uniques
type WebRTCConnection struct {
	ID uint `gorm:"primaryKey; uniqueIndex"`
	CreatedAt time.Time

	InitiatorId int `json:"initiator_id"`
	ExpirationDate int `json:"expiration_date"`
	Offer RTCSessionDescriptionInit `gorm:"embedded" json:"offer"`
	Answer RTCSessionDescriptionInit `gorm:"embedded" json:"answer"`
	// TODO cascade delete  __  constraint:OnDelete:CASCADE
	Candidates []IceCandidate `gorm:"foreignKey:ConnectionID" json:"candidates"`
}

var DB *gorm.DB

func InitDb() {
	dsn := viper.Get("DSN").(string)

	db, err := gorm.Open(postgres.Open(dsn))
	utils.ErrorHandler(err)
	DB = db
}

func Migrate() {
	DB.AutoMigrate(&WebRTCConnection{}, &IceCandidate{})
}

func CreateWebRTCConnection(con *WebRTCConnection) {
	res := DB.Create(con)
	utils.ErrorHandler(res.Error)
}

func GetActiveConnections() []WebRTCConnection {
	var peerCons []WebRTCConnection

	res := DB.Preload("Candidates").Find(&peerCons)
	utils.ErrorHandler(res.Error)

	return peerCons
}