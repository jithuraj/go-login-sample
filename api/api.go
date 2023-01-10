package api

import (
	"fmt"
	"login-sample/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
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

type user struct {
	username string
	password string
}

func listUsers(c *gin.Context) {
	db := database.OpenDB()
	defer database.CloseDB(db)
	users := database.GetAllUsers(db)

	fmt.Println(users)
	res, err := json.Marshal(users)
	if err != nil {
		panic(err.Error())
	}
	// res2 :=
	// fmt.Println(res2)
	c.String(http.StatusOK, string(res))

}
