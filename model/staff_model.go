package model

type MsStaff struct {
	NIS           string `json:"NIS" gorm:"primaryKey;size:5"`
	StaffName     string `json:"StaffName" gorm:"type:varchar(255)"`
	StaffPassword string `json:"StaffPassword" gorm:"type:varchar(255)"`
	Stsrc         string `json:"Stsrc" gorm:"size:1"`
}

type StaffLoginRequestDTO struct {
	NIS      string
	Password string
}

type StaffLoginResponseDTO struct {
	// Token        string `json:"token"`
	// RefreshToken string `json:"refreshtoken"`
	StatusCode int    `json:"statuscode"`
	Message    string `json:"message"`
	UserId     string `json:"userid"`
}
