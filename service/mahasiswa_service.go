package service

import (
	"Libeery/helper"

	"github.com/gin-gonic/gin"
)

func MhsSetupRoutes(public *gin.RouterGroup) {
	public.POST("/loginmhs",
		helper.LoginMhs,
	)
}
