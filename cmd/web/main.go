package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
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

	ticket := &Ticket{1, "Help with email", "My email is broken", 1, time.Now(), time.Now()}
	db.QueryRow("SELECT title FROM tickets WHERE id = $1", &ticket.ID)

	fmt.Println(ticket.Title)
}
