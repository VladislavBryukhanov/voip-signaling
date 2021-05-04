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
	Candidates []IceCandidate `gorm:"foreignKey:ConnectionID; constraint:OnDelete:CASCADE;" json:"candidates"`
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


func GetActiveConnections() ([]WebRTCConnection, error) {
	var peerCons []WebRTCConnection

	res := DB.Preload("Candidates").Find(&peerCons)
	return peerCons, res.Error
}

func CreateWebRTCConnection(con *WebRTCConnection) error {
	res := DB.Create(con)
	return res.Error
}

func AttachIceCandidate(ice *IceCandidate) error {
	res := DB.Create(ice)
	return res.Error
}

func DeleteConnection(connectionId uint) error {
	res := DB.Delete(&WebRTCConnection{}, connectionId)
	return res.Error
}