package model

type MsMahasiswa struct {
	NIM      string `json:"NIM" gorm:"primaryKey"`
	MhsName  string `json:"MhsName"`
	Password string `json:"Password"`
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
