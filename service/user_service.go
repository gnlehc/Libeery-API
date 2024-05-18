package service

import (
	"Libeery/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserProfileHandler(c *gin.Context) {
	userID := c.Query("userid")
	user, err := helper.GetUserProfile(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user.NIM != nil && user.NIS != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user cannot have both NIM and NIS"})
		return
	}

	var userProfile interface{}
	var profileErr error

	if user.NIM != nil {
		// Retrieve user profile from Mahasiswa table
		userProfile, profileErr = helper.GetMahasiswaProfile(c, *user.NIM)
	} else if user.NIS != nil {
		// Retrieve user profile from Staff table
		userProfile, profileErr = helper.GetStaffProfile(c, *user.NIS)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user profile not found"})
		return
	}

	if profileErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": profileErr.Error()})
		return
	}

	c.JSON(http.StatusOK, userProfile)
}

func UserRoutes(private *gin.RouterGroup) {
	private.GET("/user-profile/", UserProfileHandler)
}
