package model

import "time"

type MsAcara struct {
	AcaraID        int       `json:"AcaraID" gorm:"primaryKey;autoIncrement;size:10`
	AcaraName      string    `json:"AcaraName"`
	AcaraStartTime time.Time `json:"AcaraStartTime"`
	AcaraEndTime   time.Time `json:"AcaraEndTime"`
	AcaraDate      time.Time `json:"AcaraDate"`
	AcaraLocation  string    `json:"AcaraLocation"`
	AcaraDetails   string    `json:"AcaraDetails"`
	SpeakerName    string    `json:"SpeakerName"`
	RegisterLink   string    `json:"RegisterLink"`
	AcaraImage     string    `json:"AcaraImage""`
	Stsrc          string    `json:"Stsrc" gorm:"size:1"`
}
