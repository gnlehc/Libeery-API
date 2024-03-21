package model

import "time"

type ForLaterSession struct {
	SessionID    int       `json:"SessionID" gorm:"primaryKey"`
	StartSession time.Time `json:"StartSession"`
	EndSession   time.Time `json:"EndSession"`
}
