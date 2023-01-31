package api

import (
	"login-sample/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	db := database.OpenDB()
	defer database.CloseDB(db)
	isMatch := database.GetUserDetails(db, username, password)
	if isMatch {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "logged in"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": false, "message": "incorrect username or password"})
	}

}
