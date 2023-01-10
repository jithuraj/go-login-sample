package database

import (
	"database/sql"
	"encoding/json"
)

func GetAllUsers(db *sql.DB) string {
	users := []User{}
	rows, err := db.Query("SELECT * FROM users;")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var username string
		var password string
		err := rows.Scan(&username, &password)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, User{Username: username, Password: password})
	}
	usersJson, err := json.Marshal(users)
	if err != nil {
		panic(err.Error())
	}
	return string(usersJson)
}
