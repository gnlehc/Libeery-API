package helper

import (
	"Libeery/database"
	"Libeery/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginMhs(c *gin.Context) {
	var req model.MsMhsLoginRequestDTO
	var mhs model.MsMahasiswa

	if err := c.ShouldBindJSON(&req); err != nil {
		res := model.MsMhsLoginResponseDTO{StatusCode: 400, Message: "Invalid request", UserId: ""}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := database.GlobalDB.Where("NIM = ? AND Mhs_Password = ?", req.NIM, req.Password).First(&mhs).Error; err != nil {
		res := model.MsMhsLoginResponseDTO{StatusCode: 401, Message: "Credentials not matched", UserId: ""}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	var userid string
	if err := database.GlobalDB.Model(&model.MsUser{}).Where("nim = ?", req.NIM).Pluck("user_id", &userid).Error; err != nil {
		res := model.MsMhsLoginResponseDTO{StatusCode: 501, Message: "Not Implemented", UserId: ""}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := model.MsMhsLoginResponseDTO{StatusCode: 200, Message: "Success", UserId: userid}
	c.JSON(http.StatusOK, res)
}
