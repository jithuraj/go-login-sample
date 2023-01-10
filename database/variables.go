package database

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
