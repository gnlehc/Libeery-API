package service

import (
	"Libeery/helper"

	"github.com/gin-gonic/gin"
)

func StaffRoutes(public *gin.RouterGroup) {
	public.POST("/loginstaff",
		helper.LoginStaff,
	)
}
