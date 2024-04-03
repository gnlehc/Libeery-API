package service

import (
	"Libeery/helper"

	"github.com/gin-gonic/gin"
)

func MhsRoutes(public *gin.RouterGroup) {
	public.POST("/loginmhs",
		helper.LoginMhs,
	)
}
