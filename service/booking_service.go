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

	if err := helper.CreateBookingForLater(c, req); err != nil {
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

func BookASessionForNowHandler(c *gin.Context) {
	var req model.BookingRequestForNowDTO

	// Parse JSON body into BookingRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, output.BaseOutput{
			Message:    err.Error(),
			StatusCode: 403,
		})
		return
	}

	if err := helper.CreateBookingForNow(c, req); err != nil {
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

func GetUserBooking(c *gin.Context) {
	userID := c.Query("userID")
	resultList, err := helper.GetUserBookingData(c, userID)
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

func CheckInBookingHandler(c *gin.Context) {
	var req model.CheckInOutBookingRequestDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, output.BaseOutput{
			Message:    err.Error(),
			StatusCode: 403,
		})
		return
	}

	if err := helper.CheckInBooking(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, output.BaseOutput{
			Message:    err.Error(),
			StatusCode: 403,
		})
		return
	}

	c.JSON(http.StatusOK, output.BaseOutput{
		Message:    "Check In Successful",
		StatusCode: 200,
	})
}

func CheckOutBookingHandler(c *gin.Context) {
	var req model.CheckInOutBookingRequestDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, output.BaseOutput{
			Message:    err.Error(),
			StatusCode: 403,
		})
		return
	}

	if err := helper.CheckOutBooking(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, output.BaseOutput{
			Message:    err.Error(),
			StatusCode: 403,
		})
		return
	}

	c.JSON(http.StatusOK, output.BaseOutput{
		Message:    "Check Out Successful",
		StatusCode: 200,
	})
}

func CheckBookingStatusHandler(c *gin.Context) {
	var req model.CheckBookingStatusRequestDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	isCheckedIn, err := helper.GetCheckInStatus(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get booking status",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isCheckedIn": isCheckedIn,
	})
}
func DeleteExpiredBookingsHandler(c *gin.Context) {
	bookingId := c.Query("id")
	if err := helper.DeleteExpiredBookings(c, bookingId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete expired bookings",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Expired bookings deleted successfully",
	})
}

func GetBookingByIDHandler(c *gin.Context) {
	bookingID := c.Query("id")

	booking, err := helper.GetBookingDataByID(c, bookingID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	c.JSON(http.StatusOK, booking)
}

func BookingRoutes(private *gin.RouterGroup) {
	private.GET("/bookings", BookingHandler)
	private.POST("/bookSession", BookASessionHandler)
	private.POST("/bookSession-forNow", BookASessionForNowHandler)
	private.GET("/user-bookings/", GetUserBooking)
	private.POST("/check-in", CheckInBookingHandler)
	private.POST("/check-out", CheckOutBookingHandler)
	private.POST("/check-booking-status", CheckBookingStatusHandler)
	private.POST("/get-booking-by-id", GetBookingByIDHandler)
	private.POST("/delete-expired-booking", DeleteExpiredBookingsHandler)
}
