package helper

import (
	"Libeery/database"
	model "Libeery/model"
	"Libeery/output"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context, userID string) (*model.MsUser, error) {
	db := database.GlobalDB
	var user model.MsUser
	if err := db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}

func GetMahasiswaProfile(c *gin.Context, nim string) (*output.UserProfileMahasiswaOutput, error) {
	db := database.GlobalDB
	var mahasiswa model.MsMahasiswa
	if err := db.First(&mahasiswa, "nim = ?", nim).Error; err != nil {
		return nil, err
	}
	userProfile := &output.UserProfileMahasiswaOutput{
		NIM:     mahasiswa.NIM,
		MhsName: mahasiswa.MhsName,
	}
	return userProfile, nil
}

func GetStaffProfile(c *gin.Context, nis string) (*output.UserProfileStaffOutput, error) {
	db := database.GlobalDB
	var staff model.MsStaff
	if err := db.First(&staff, "nis = ?", nis).Error; err != nil {
		return nil, err
	}
	userProfile := &output.UserProfileStaffOutput{
		NIS:       staff.NIS,
		StaffName: staff.StaffName,
	}
	return userProfile, nil
}
