package service

import (
	"Libeery/helper"
	"Libeery/model"
	"Libeery/output"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BookASessionHandler(c *gin.Context) {
	var req model.BookingRequestDTO

	// Parse JSON body into BookingRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := helper.CreateBooking(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output.BaseOutput{
		Message:    "Booking Successfull",
		StatusCode: 200,
	})
}

func BookingRoutes(public *gin.RouterGroup) {
	public.POST("/bookSession", BookASessionHandler)
}
