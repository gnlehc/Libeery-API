package model

type MsMahasiswa struct {
	NIM         string `json:"NIM" gorm:"primaryKey;size:10"`
	MhsName     string `json:"MhsName" gorm:"type:varchar(255)"`
	MhsPassword string `json:"MhsPassword" gorm:"type:varchar(255)"`
	Stsrc       string `json:"Stsrc" gorm:"size:1"`
}

type LoginRequestDTO struct {
	NIM      string
	Password string
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
	StatusCode   string `json:"statuscode"`
	Message      string `json:"message"`
}
