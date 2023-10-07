package database

import (
	"database/sql"
	"log"
	"os"
	"sync"
)

func getDB() *Queries {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	DB := New(db)

	if err != nil {
		log.Fatal(err)
	}

	return DB
}

var GetDB = sync.OnceValue(getDB)
