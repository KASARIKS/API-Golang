package main

import (
	"log"

	"github.com/kasariks/api_golang/cmd/api"
	"github.com/kasariks/api_golang/config"
	"github.com/kasariks/api_golang/db"
)

func main() {
	db, err := db.NewMySQLStorage(config.Envs.DBName)
	if err != nil {
		log.Fatal(err)
	}

	// initStorage(db)

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// func initStorage(db *sql.DB) {
// 	err := db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("DB has been successfully connected.")
// }
