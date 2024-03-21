package model

type MsBookingStatus struct {
	BookingStatusID int    `json:"BookingStatusID" gorm:"primaryKey"`
	StatusTitle     string `json:"StatusTitle" gorm:"type:varchar(255)"`
	Stsrc           string `gorm:"size:1" json:"Stsrc"`
}
