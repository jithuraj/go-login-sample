package database

import (
	"database/sql"
)

func GetUserDetails(db *sql.DB, username string, password string) bool {
	rows, err := db.Query("SELECT * FROM users WHERE username=? AND password=?", username, password)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	if rows.Next() {
		// match
		return true
	} else {
		// not a match
		return false
	}
}
