package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Welcome To Libeery",
		})
	})

	api := r.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/login")
			public.POST("/register")
		}
	}
	return r
}
