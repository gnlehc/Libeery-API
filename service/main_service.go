package service

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		public := api.Group("/public")
		{
			MhsSetupRoutes(public)
			StaffSetupRoutes(public)
		}
	}
}
