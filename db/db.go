package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func NewMySQLStorage(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbName)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
