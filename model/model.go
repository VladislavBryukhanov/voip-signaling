package model

import (
	"log"
	"time"
	"gorm.io/gorm"
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

	InitiatorId int `json:"initiator_id"`
	ExpirationDate int `json:"expiration_date"`
	Offer RTCSessionDescriptionInit `gorm:"embedded" json:"offer"`
	Answer RTCSessionDescriptionInit `gorm:"embedded" json:"answer"`
	// Candidates []IceCandidate `gorm:"embedded" json:"cadidates"`
}

var DB *gorm.DB

func SetDatabase(db *gorm.DB, sync bool) {
	DB = db

	if sync {
		db.AutoMigrate(&WebRTCConnection{})
	}
}

func CreateWebRTCConnection(con *WebRTCConnection) {
	res := DB.Create(con)

	if res.Error != nil {
		log.Fatal(res.Error)
	}
}

func GetActiveConnections() []WebRTCConnection {
	var peerCons []WebRTCConnection
	res := DB.Find(&peerCons)

	if res.Error != nil {
		log.Fatal(res.Error)
	}

	return peerCons
}