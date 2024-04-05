package helper

import (
	"Libeery/database"
	"Libeery/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginStaff(c *gin.Context) {
	var req model.MsStaffLoginRequestDTO
	var staff model.MsStaff

	if err := c.ShouldBindJSON(&req); err != nil {
		res := model.MsStaffLoginResponseDTO{StatusCode: 400, Message: "Invalid request", UserId: ""}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := database.GlobalDB.Where("nis = ? AND staff_password = ?", req.NIS, req.Password).First(&staff).Error; err != nil {
		res := model.MsStaffLoginResponseDTO{StatusCode: 401, Message: "Credentials not matched", UserId: ""}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var userid string
	if err := database.GlobalDB.Model(&model.MsUser{}).Where("nis = ?", req.NIS).Pluck("user_id", &userid).Error; err != nil {
		res := model.MsStaffLoginResponseDTO{StatusCode: 501, Message: "Not Implemented", UserId: ""}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := model.MsStaffLoginResponseDTO{StatusCode: 200, Message: "Success", UserId: userid}
	c.JSON(http.StatusOK, res)
}
