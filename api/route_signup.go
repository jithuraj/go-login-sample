package api

import (
	"login-sample/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	db := database.OpenDB()
	defer database.CloseDB(db)
	var status bool = database.Insert(db, username, password)
	if status {
		c.JSON(http.StatusOK, gin.H{
			"status": true, "message": "signup successful",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": false, "message": "something went wrong!!!",
		})
	}
}
