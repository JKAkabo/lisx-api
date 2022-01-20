package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sqlx.DB

func Connect() *sqlx.DB {
	log.Println(os.Getenv("DATABASE_URL"))
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return db
}
