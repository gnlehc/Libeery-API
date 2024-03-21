package model

import (
	"time"

	"github.com/google/uuid"
)

type TrBooking struct {
	BookingID       uuid.UUID `json:"BookingID" gorm:"primaryKey;unique;default:gen_random_uuid()"`
	UserID          int       `json:"UserID" gorm:"foreignKey:UserID"`
	BookingStatusID int       `json:"BookingStatusID" gorm:"foreignKey:BookingStatusID"`
	LokerID         int       `json:"LokerID" gorm:"foreignKey:LokerID"`
	StartSession    time.Time `json:"StartSession"`
	EndSession      time.Time `json:"EndSession"`
	CheckInTime     time.Time `json:"CheckInTime"`
	CheckOutTime    time.Time `json:"CheckOutTime"`
	CreatedAt       time.Time `json:"CreatedAt"`
	UpdatedAt       time.Time `json:"UpdatedAt"`
	Stsrc           string    `json:"Stsrc" gorm:"size:1"`
}
