package helper

import (
	"Libeery/database"
	"Libeery/model"

	"github.com/gin-gonic/gin"
)

func GetAllSessionData(c *gin.Context) ([]*model.MsSession, error) {
	db := database.GlobalDB

	var sessions []*model.MsSession
	if err := db.Find(&sessions).Error; err != nil {
		return nil, err
	}

	return sessions, nil
}
