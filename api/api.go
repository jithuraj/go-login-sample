package api

import (
	"login-sample/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Api() {
	r := gin.Default()
	r.GET("/", Root)
	r.POST("/signup", Signup)
	r.Run()

}

func Root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"welcome to ooty": "nice to meet you",
	})
}

func Signup(c *gin.Context) {
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
