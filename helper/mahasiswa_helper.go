package helper

import (
	"Libeery/database"
	"Libeery/model"
	"Libeery/output"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginMhs(c *gin.Context) {
	var req model.MhsLoginRequestDTO
	var mhs model.MsMahasiswa

	if err := c.ShouldBindJSON(&req); err != nil {
		res := output.LoginResponseDTO{StatusCode: 400, Message: "Invalid request", UserId: ""}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := database.GlobalDB.Where("NIM = ? AND Mhs_Password = ?", req.NIM, req.Password).First(&mhs).Error; err != nil {
		res := output.LoginResponseDTO{StatusCode: 401, Message: "Credentials not matched", UserId: ""}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	var userid string
	if err := database.GlobalDB.Model(&model.MsUser{}).Where("nim = ?", req.NIM).Pluck("user_id", &userid).Error; err != nil {
		res := output.LoginResponseDTO{StatusCode: 501, Message: "Not Implemented", UserId: ""}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	var username string
	if err := database.GlobalDB.Model(&model.MsMahasiswa{}).Where("NIM = ?", req.NIM).Pluck("mhs_name", &username).Error; err != nil {
		res := output.LoginResponseDTO{StatusCode: 501, Message: "Not Implemented", UserId: ""}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := output.LoginResponseDTO{StatusCode: 200, Message: "Success", UserId: userid, Username: username}
	c.JSON(http.StatusOK, res)
}
