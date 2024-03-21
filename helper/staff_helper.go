package helper

import (
	"Libeery/database"
	"Libeery/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginStaff(c *gin.Context) {
	var req model.StaffLoginRequestDTO
	var staff model.MsStaff

	if err := c.ShouldBindJSON(&req); err != nil {
		res := model.StaffLoginResponse{StatusCode: "400", Message: "Invalid request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := database.GlobalDB.Where("nis = ?", req.NIS).First(&staff).Error; err != nil {
		res := model.StaffLoginResponse{StatusCode: "401", Message: "NIS not found"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := database.GlobalDB.Where("nis = ? AND staff_password = ?", req.NIS, req.Password).First(&staff).Error; err != nil {
		res := model.StaffLoginResponse{StatusCode: "401", Message: "Password not matched!"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := database.GlobalDB.Model(&staff).Update("Stsrc", "A").Error; err != nil {
		res := model.StaffLoginResponse{StatusCode: "500", Message: "Failed to update Stsrc attribute"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := model.StaffLoginResponse{StatusCode: "200", Message: "Success", Staff: staff}
	c.JSON(http.StatusOK, res)
}
