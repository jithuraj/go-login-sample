package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func OpenDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Ahj&&65tttttttt@/login")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}

func CloseDB(db *sql.DB) {
	db.Close()
}

func Insert(db *sql.DB, username string, password string) {
	db.Exec("INSERT INTO users (username,password) values (?,?);", username, password)
}

func GetAllUsers(db *sql.DB) []User {
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
	return users
}
