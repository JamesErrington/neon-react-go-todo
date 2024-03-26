package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	connStr, found := os.LookupEnv("DB_CONN_STRING")
	if !found {
		log.Fatal("No value found for DB_CONN_STRING")
	}

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database!")
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database closed!")
}
