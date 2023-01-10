package database

import "database/sql"

func CloseDB(db *sql.DB) {
	db.Close()
}
