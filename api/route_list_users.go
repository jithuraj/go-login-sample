package api

import (
	"login-sample/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func listUsers(c *gin.Context) {
	db := database.OpenDB()
	defer database.CloseDB(db)
	users := database.GetAllUsers(db)
	c.String(http.StatusOK, users)
}
