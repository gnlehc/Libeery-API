package output

type UserProfileMahasiswaOutput struct {
	NIM     string `gorm:"size:10" json:"NIM,omitempty"`
	MhsName string `json:"MhsName" gorm:"type:varchar(255)"`
}

type UserProfileStaffOutput struct {
	NIS       string `json:"NIS" gorm:"primaryKey;size:5"`
	StaffName string `json:"StaffName" gorm:"type:varchar(255)"`
}
