package service

import (
	"Libeery/helper"
	"Libeery/output"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SessionHandler(c *gin.Context) {
	resultList, err := helper.GetAllSessionData(c)
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
	c.JSON(http.StatusOK, output.SessionOutput{
		Data: resultList,
		BaseOutput: output.BaseOutput{
			Message:    "Success",
			StatusCode: http.StatusOK,
		},
	})
}

func SessionRoutes(private *gin.RouterGroup) {
	private.GET("/sessions", SessionHandler)
}
