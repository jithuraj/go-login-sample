package database

import "database/sql"

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
