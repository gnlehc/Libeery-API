// service package
package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		public := api.Group("/public")
		{
			MhsRoutes(public)
			StaffRoutes(public)
		}
		private := api.Group("/private")
		{
			BookingRoutes(private)
			SessionRoutes(private)
			LokerRoutes(private)
			AcaraRoutes(private)
		}
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Libeery API!",
		})
	})
}
