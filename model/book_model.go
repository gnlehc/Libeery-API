package model

type MsBook struct {
	BookID    int    `json:"BookID" gorm:"primaryKey;autoIncrement:true;not null"`
	ISBN      string `json:"ISBN" gorm:"type:char(14);unique"`
	Title     string `json:"Title"`
	Author    string `json:"Author"`
	Publisher string `json:"Publisher"`
	Edition   string `json:"Edition"`
	Year      int    `json:"Year"`
	Abstract  string `json:"Abstract"`
	Stock     int    `json:"Stock"`
	Stsrc     string `json:"Stsrc"`
}
