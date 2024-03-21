package model

type MsLoker struct {
	LockerID     int `gorm:"primaryKey" json:"LockerID"`
	RowNumber    int `json:"RowNumber"`
	ColumnNumber int `json:"ColumnNumber"`
}
