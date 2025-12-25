package db

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func NewMySQLStorage(dbName string) (*sql.DB, error) {
	_, err := os.Stat(dbName)
	var install bool

	if err != nil {
		install = true
	}

	db, err := sql.Open("sqlite", dbName)
	if err != nil {
		log.Fatal(err)
	}

	if install {
		if err := createTables(db); err != nil {
			return nil, err
		}
	}

	return db, nil
}
