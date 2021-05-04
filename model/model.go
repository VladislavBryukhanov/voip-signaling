package model

import (
	"time"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/VladislavBryukhanov/voip-signaling/utils"
)

type RTCSessionDescription struct {
	ID uint `gorm:"primaryKey; uniqueIndex" json:"-"`
	ConnectionID uint `gorm:"uniqueIndex:idx_type_con_id" json:"-"`

	Sdp string `json:"sdp"`
	Type string `gorm:"type:session_description_type; uniqueIndex:idx_type_con_id" json:"type"`
}

type IceCandidate struct {
	ID uint `gorm:"primaryKey; uniqueIndex" json:"-"`
	ConnectionID uint `json:"-"`

	Candidate string `json:"candidate"`
	SdpMLineIndex int `json:"sdpMLineIndex"`
	SdpMid int `json:"sdpMid"`
}

type WebRTCConnection struct {
	ID uint `gorm:"primaryKey; uniqueIndex"`
	CreatedAt time.Time `json:"-"`

	InitiatorId int `json:"initiator_id"`
	ExpirationDate int `json:"expiration_date"`
	SessionDescriptions []RTCSessionDescription `gorm:"foreignKey:ConnectionID; constraint:OnDelete:CASCADE" json:"session_descriptions"`
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
	DB.Exec(`
		DO $$ BEGIN
			CREATE TYPE session_description_type AS ENUM('offer', 'answer');
		EXCEPTION
			WHEN duplicate_object THEN null;
		END $$;
	`)
	DB.AutoMigrate(&WebRTCConnection{}, &IceCandidate{}, &RTCSessionDescription{})
}

func GetActiveConnections() ([]WebRTCConnection, error) {
	var peerCons []WebRTCConnection

	res := DB.Preload("Candidates").Preload("SessionDescriptions").Find(&peerCons)
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

func AttachSessionDescription(session *RTCSessionDescription) error {
	res := DB.Create(session)
	return res.Error
}

func DeleteConnection(connectionId uint) error {
	res := DB.Delete(&WebRTCConnection{}, connectionId)
	return res.Error
}