package database

import "database/sql"

func Insert(db *sql.DB, username string, password string)bool {
	_,err :=db.Exec("INSERT INTO users (username,password) values (?,?);", username, password)
	if err!=nil {
		return false
	}else{
		return true
	}
}
