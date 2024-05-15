package service

import (
	"Libeery/helper"
	"Libeery/model"
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

func GetAcaraDetailsHandler(c *gin.Context) {
	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}
	acara, err := helper.GetAcaraDetails(c, idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}
	c.JSON(http.StatusOK, output.GetAcaraDetailsOutput{
		Data: acara,
		BaseOutput: output.BaseOutput{
			Message:    "Success",
			StatusCode: 200,
		},
	})
}

func CreateAcaraHandler(c *gin.Context) {
	var reqBody model.MsAcaraRequestDTO

	// Bind the request body to the request DTO
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the helper function to create the Acara
	err := helper.CreateAcara(c, reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Acara"})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, output.BaseOutput{
		Message:    "Created Acara Successfully",
		StatusCode: http.StatusOK,
	})
}

func AcaraRoutes(private *gin.RouterGroup) {
	private.GET("/acara", GetAcaraPaginationHandler)
	private.GET("/get-acara", GetAllAcaraHandler)
	private.GET("/acara-detail/", GetAcaraDetailsHandler)
	private.POST("/create-acara", CreateAcaraHandler)
}
