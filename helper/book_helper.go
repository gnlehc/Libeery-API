package helper

import (
	"Libeery/database"
	"Libeery/model"

	"github.com/gin-gonic/gin"
)

func GetListOfBooks(c *gin.Context) ([]*model.MsBook, error) {
	db := database.GlobalDB
	var ListOfBooks []*model.MsBook
	if err := db.Find(&ListOfBooks).Error; err != nil {
		return nil, err
	}

	return ListOfBooks, nil
}
