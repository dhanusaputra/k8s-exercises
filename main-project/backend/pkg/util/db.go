package util

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

// InitDB ...
func InitDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
		"postgres-svc",
		"5432",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"))

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS todo (
    id SERIAL PRIMARY KEY, 
    text VARCHAR(255) NOT NULL,
    done BOOLEAN NOT NULL
  );`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("successfully init db")

	return db
}

// PingDB ...
func PingDB() error {
	if db == nil {
		return errors.New("db is nil")
	}
	return db.Ping()
}
