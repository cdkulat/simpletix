package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	_, err := godotenv.Load()
	if err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	connStr, exists := os.LookupEnv("CONNECTION_STRING")
	if !exists {
		log.Print("Connection string not found...")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	id := 1
	rows, err := db.Query("SELECT title FROM tickets WHERE id = $1", id)

	fmt.Println(rows)
}
