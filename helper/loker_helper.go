package helper

import (
	"Libeery/database"
	"Libeery/model"

	"github.com/gin-gonic/gin"
)

func GetAllLokerData(c *gin.Context) ([]*model.MsLoker, error) {
	db := database.GlobalDB

	var lokers []*model.MsLoker
	if err := db.Find(&lokers).Error; err != nil {
		return nil, err
	}

	return lokers, nil
}

func GetAllLokerDataBySessionID(c *gin.Context, sessionID int) ([]*model.MsLoker, error) {
	db := database.GlobalDB

	var lokers []*model.MsLoker

	err := db.Raw("SELECT * FROM ms_lokers WHERE availability = 'Active' AND locker_id NOT IN (SELECT DISTINCT loker_id FROM tr_bookings WHERE session_id = ?)", sessionID).Scan(&lokers).Error
	if err != nil {
		return nil, err
	}

	return lokers, nil
}
