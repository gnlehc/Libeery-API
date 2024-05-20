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

func GetBookDetails(c *gin.Context, id int) (*model.MsBook, error) {
	db := database.GlobalDB
	var book model.MsBook
	if err := db.Where("book_id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}
