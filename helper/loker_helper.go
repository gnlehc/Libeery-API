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

	err := db.Raw("SELECT * FROM ms_lokers WHERE availability = 'Active' AND loker_id NOT IN (SELECT DISTINCT loker_id FROM tr_bookings WHERE session_id = ?)", sessionID).Scan(&lokers).Error
	if err != nil {
		return nil, err
	}

	return lokers, nil
}

func GetLokerDataByMultipleSessionID(c *gin.Context, sessionIDs []int) ([]*model.MsLoker, error) {
	db := database.GlobalDB

	var lokers []*model.MsLoker

	// Construct the query to exclude bookings for the given session IDs
	query := db.Table("ms_lokers").
		Where("availability = ?", "Active").
		Where("loker_id NOT IN (?)",
			db.Table("tr_bookings").Select("DISTINCT loker_id").Where("session_id IN (?)", sessionIDs))

	// Execute the query
	err := query.Scan(&lokers).Error
	if err != nil {
		return nil, err
	}

	return lokers, nil
}
