package service

import (
	"Libeery/helper"
	"Libeery/model"
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

	// Call MakeBooking function from helper package
	if err := helper.MakeBooking(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking successful"})
}

func BookingRoutes(public *gin.RouterGroup) {
	public.POST("/book_a_session", BookASessionHandler)
}
