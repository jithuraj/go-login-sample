package database

import "database/sql"

func Insert(db *sql.DB, username string, password string) {
	db.Exec("INSERT INTO users (username,password) values (?,?);", username, password)
}
