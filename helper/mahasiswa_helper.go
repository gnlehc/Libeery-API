package helper

import (
	"Libeery/database"
	"Libeery/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginMhs(c *gin.Context) {
	var req model.MhsLoginRequestDTO
	var mhs model.MsMahasiswa

	if err := c.ShouldBindJSON(&req); err != nil {
		res := model.MhsLoginResponse{StatusCode: "400", Message: "Invalid request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := database.GlobalDB.Where("NIM = ?", req.NIM).First(&mhs).Error; err != nil {
		res := model.MhsLoginResponse{StatusCode: "401", Message: "NIM not found"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := database.GlobalDB.Where("NIM = ? AND Mhs_Password = ?", req.NIM, req.Password).First(&mhs).Error; err != nil {
		res := model.MhsLoginResponse{StatusCode: "401", Message: "Password not matched!"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := database.GlobalDB.Model(&mhs).Update("Stsrc", "A").Error; err != nil {
		res := model.MhsLoginResponse{StatusCode: "500", Message: "Failed to update Stsrc attribute"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := model.MhsLoginResponse{StatusCode: "200", Message: "Success", Mhs: mhs}
	c.JSON(http.StatusOK, res)
}
