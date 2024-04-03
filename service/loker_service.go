package service

import (
	"Libeery/helper"
	"Libeery/output"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LokerHandler(c *gin.Context) {
	resultList, err := helper.GetAllLokerData(c)
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
	c.JSON(http.StatusOK, output.LokerOutput{
		Data: resultList,
		BaseOutput: output.BaseOutput{
			Message:    "Success",
			StatusCode: http.StatusOK,
		},
	})
}
func LockerBySessionIDDataHandler(c *gin.Context) {
	sessionID := c.Query("session_id")
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, output.BaseOutput{
			Message:    "Missing session_id parameter",
			StatusCode: 400,
		})
		return
	}

	// Convert sessionID to integer
	sessionIDInt, err := strconv.Atoi(sessionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, output.BaseOutput{
			Message:    "Invalid session_id parameter",
			StatusCode: 400,
		})
		return
	}

	// Get locker data for the specified session ID
	lockerData, err := helper.GetAllLokerDataBySessionID(c, sessionIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, output.BaseOutput{
			Message:    err.Error(),
			StatusCode: 500,
		})
		return
	}

	// Return locker data as JSON response
	c.JSON(http.StatusOK, output.LokerOutput{
		Data: lockerData,
		BaseOutput: output.BaseOutput{
			Message:    "Success",
			StatusCode: 200,
		},
	})
}
func LokerRoutes(private *gin.RouterGroup) {
	private.GET("/lokers", LokerHandler)
	private.GET("/lokerBySessionID", LockerBySessionIDDataHandler)
}
