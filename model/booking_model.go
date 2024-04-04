package model

import (
	"time"

	"github.com/google/uuid"
)

type TrBooking struct {
	BookingID       uuid.UUID `json:"BookingID" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID          uuid.UUID `json:"UserID"`
	SessionID       int       `json:"SessionID"`
	BookingStatusID int       `json:"BookingStatusID"`
	LokerID         int       `json:"LokerID"`
	CreatedAt       time.Time `json:"CreatedAt"`
	UpdatedAt       time.Time `json:"UpdatedAt"`
	Stsrc           string    `json:"Stsrc" gorm:"size:1"`
}

type BookingRequestDTO struct {
	UserID    uuid.UUID `json:"UserID"`
	SessionID int       `json:"SessionID"`
	LokerID   int       `json:"LokerID"`
}

type CheckBookingRequestDTO struct {
	UserID    uuid.UUID `json:"UserID"`
	SessionID int       `json:"SessionID"`
}

// type CheckInRequestDTO struct {
// 	UserID uuid.UUID `json:"UserID"`
// 	BookingStatusID int `json:"BookingStatusID"`
// 	LokerID int `json:"BookingStatusID"`
// }
