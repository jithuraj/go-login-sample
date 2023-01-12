package api

import (
	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()
	r.GET("/", root)
	r.POST("/signup", signup) //username,password
	r.POST("/login",login) //username.password
	r.GET("/list", listUsers)
	r.Run()
}
