package api

import (
	"login-sample/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	db := database.OpenDB()
	defer database.CloseDB(db)
	isMatch := database.GetUserDetails(db, username, password)
	if isMatch {
		c.JSON(http.StatusOK, gin.H{"status": "user found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "user not found"})
	}

}
