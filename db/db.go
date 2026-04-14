package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Db() *sql.DB {
	if db == nil {
		connect()
		return db
	}
	return db
}

func connect() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	sslMode := os.Getenv("DB_SSL_MODE")
	connStr := fmt.Sprintf("user=%s password=%s database=%s port=%d sslmode=%s", user, password, database, port, sslMode)
	if db, err = sql.Open("postgres", connStr); err != nil {
		log.Fatal(err)
	}
}
