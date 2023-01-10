package api

import (
	"login-sample/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"password": password,
	})

	db := database.OpenDB()
	defer database.CloseDB(db)
	database.Insert(db, username, password)
}
