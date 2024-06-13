package service

import (
	"Libeery/helper"
	"Libeery/output"
	"net/http"
	"strconv"
	"strings"

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
func LokerBySessionIDDataHandler(c *gin.Context) {
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
	lokerData, err := helper.GetAllLokerDataBySessionID(c, sessionIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, output.BaseOutput{
			Message:    err.Error(),
			StatusCode: 500,
		})
		return
	}

	// Return locker data as JSON response
	c.JSON(http.StatusOK, output.LokerOutput{
		Data: lokerData,
		BaseOutput: output.BaseOutput{
			Message:    "Success",
			StatusCode: 200,
		},
	})
}

func LokerByMultipleSessionIDDataHandler(c *gin.Context) {
	sessionIDsParam := c.Query("session_ids")
	if sessionIDsParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Missing session_ids parameter",
			"statusCode": 400,
		})
		return
	}

	// Split sessionIDsParam into a slice of strings
	sessionIDsStr := strings.Split(sessionIDsParam, ",")

	// Convert sessionIDsStr to a slice of integers
	sessionIDs := make([]int, len(sessionIDsStr))
	for i, idStr := range sessionIDsStr {
		id, err := strconv.Atoi(strings.TrimSpace(idStr))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message":    "Invalid session_ids parameter",
				"statusCode": 400,
			})
			return
		}
		sessionIDs[i] = id
	}

	// Get locker data for the specified session IDs
	lokerData, err := helper.GetLokerDataByMultipleSessionID(c, sessionIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":    err.Error(),
			"statusCode": 500,
		})
		return
	}

	// Return locker data as JSON response
	c.JSON(http.StatusOK, gin.H{
		"data": lokerData,
		"baseOutput": gin.H{
			"message":    "Success",
			"statusCode": 200,
		},
	})
}
func LokerRoutes(private *gin.RouterGroup) {
	private.GET("/lokers", LokerHandler)
	private.GET("/lokerBySessionID", LokerBySessionIDDataHandler)
	private.GET("/LokerByMultipleSessionID", LokerByMultipleSessionIDDataHandler)
}
