package db

import (
	"database/sql"
	"log"
	_"github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() *sql.DB{
	var err error
	db, err = sql.Open("sqlite3", "./files.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS files (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		data BLOB NOT NULL,
		name TEXT
	);`)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
	return db
}
