package model

type MsLoker struct {
	LokerID      int    `gorm:"primaryKey" json:"LokerID"`
	RowNumber    int    `json:"RowNumber"`
	ColumnNumber int    `json:"ColumnNumber"`
	Availability string `json:"Availability"`
	Stsrc        string `json:"Stsrc" gorm:"size:1"`
}
