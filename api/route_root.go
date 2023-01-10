package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"welcome to ooty": "nice to meet you",
	})
}
