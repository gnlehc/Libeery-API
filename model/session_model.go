package model

import "time"

type MsSession struct {
	SessionID    int       `json:"SessionID" gorm:"primaryKey"`
	StartSession time.Time `json:"StartSession"`
	EndSession   time.Time `json:"EndSession"`
}

// type ListMsSession struct {
// 	ListMsSession []*MsSession `gorm:"-" json:"sessions"`
// }
