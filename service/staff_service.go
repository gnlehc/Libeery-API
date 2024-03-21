package service

import (
	"Libeery/helper"

	"github.com/gin-gonic/gin"
)

func StaffSetupRoutes(public *gin.RouterGroup) {
	public.POST("/loginstaff",
		helper.LoginStaff,
	)
}
