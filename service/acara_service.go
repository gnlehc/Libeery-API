package service

import (
	"Libeery/helper"
	"Libeery/output"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllAcaraHandler(c *gin.Context) {
	acara, err := helper.GetAllAcara(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}
	c.JSON(http.StatusOK, output.GetAcaraOutput{
		Data: acara,
		BaseOutput: output.BaseOutput{
			Message:    "Success",
			StatusCode: 200,
		},
	})
}
func GetAcaraPaginationHandler(c *gin.Context) {
	page := c.DefaultQuery("page", "1")  // Default to page 1 if not provided
	take := c.DefaultQuery("take", "10") // Default to taking 10 items per page if not provided
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}
	takeInt, err := strconv.Atoi(take)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid take parameter"})
		return
	}

	acara, err := helper.GetAcaraPagination(c, pageInt, takeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}
	c.JSON(http.StatusOK, output.GetAcaraOutput{
		Data: acara,
		BaseOutput: output.BaseOutput{
			Message:    "Success",
			StatusCode: 200,
		},
	})
}

func AcaraRoutes(private *gin.RouterGroup) {
	private.GET("/acara", GetAcaraPaginationHandler)
	private.GET("/get-acara", GetAllAcaraHandler)
}
