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
		c.JSON(http.StatusBadRequest, output.BaseOutput{
			Message:    err.Error(),
			StatusCode: 403,
		})
		return
	}

	if err := helper.CreateBooking(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, output.BaseOutput{
			Message:    err.Error(),
			StatusCode: 403,
		})
		return
	}

	c.JSON(http.StatusOK, output.BaseOutput{
		Message:    "Booking Successfull",
		StatusCode: 200,
	})
}

func BookingHandler(c *gin.Context) {
	resultList, err := helper.GetAllBookingData(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, output.SessionOutput{
			Data: nil,
			BaseOutput: output.BaseOutput{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			},
		})
		return
	}
	c.JSON(http.StatusOK, output.BookingOutput{
		Data: resultList,
		BaseOutput: output.BaseOutput{
			Message:    "Success",
			StatusCode: http.StatusOK,
		},
	})
}

func BookingRoutes(private *gin.RouterGroup) {
	private.GET("/bookings", BookingHandler)
	private.POST("/bookSession", BookASessionHandler)
}