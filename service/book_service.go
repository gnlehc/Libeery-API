package service

import (
	"Libeery/helper"
	"Libeery/output"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListBookHandler(c *gin.Context) {
	ListOfBooks, err := helper.GetListOfBooks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}
	c.JSON(http.StatusOK, output.ListBookOutput{
		Data: ListOfBooks,
		BaseOutput: output.BaseOutput{
			Message:    "Success",
			StatusCode: 200,
		},
	})
}

func BookRoutes(private *gin.RouterGroup) {
	private.GET("/get-all-books", GetListBookHandler)
}
