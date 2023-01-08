package api

import (
	"fmt"
	"login-sample/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Api() {
	r := gin.Default()
	r.GET("/", root)
	r.POST("/signup", signup)
	r.GET("/list", listUsers)
	r.Run()

}

func root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"welcome to ooty": "nice to meet you",
	})
}

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

func listUsers(c *gin.Context) {
	db := database.OpenDB()
	defer database.CloseDB(db)
	users := database.GetAllUsers(db)

	fmt.Println(users)

	c.JSON(http.StatusOK, gin.H{
		"username": users[0],
	})

}
