package service

import (
	"Libeery/helper"
	"Libeery/output"
	"net/http"
	"strconv"

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

func GetBookDetailsHandler(c *gin.Context) {
	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	book, err := helper.GetBookDetails(c, idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}
	c.JSON(http.StatusOK, output.GetBookDetailsOutput{
		Data: book,
		BaseOutput: output.BaseOutput{
			Message:    "Success",
			StatusCode: 200,
		},
	})
}

func BookRoutes(private *gin.RouterGroup) {
	private.GET("/get-all-books", GetListBookHandler)
	private.GET("/book-detail/", GetBookDetailsHandler)
}
