package driver

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {

	db, err := sql.Open("postgres", os.Getenv("CONNECTION_STRING"))

	if err != nil {
		log.Panic(err)
	}
	return db
}
