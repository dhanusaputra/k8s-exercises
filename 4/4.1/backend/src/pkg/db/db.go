package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
  "errors"
)

var db *sql.DB

// Init ...
func Init() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"))

  var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS todo (
    id SERIAL PRIMARY KEY, 
    text VARCHAR(255) NOT NULL
  );`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("successfully init db")

	return db
}

// Ping ...
func Ping() error {
  if db == nil {
    return errors.New("db is nil")
  }
	return db.Ping()
}
