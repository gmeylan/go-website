package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// _, err = db.Exec(`
	// CREATE TABLE IF NOT EXISTS items (
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	name TEXT NOT NULL
	// )
	// `)
	// if err != nil {
	// 	return nil, err
	// }

	return db, nil
}
