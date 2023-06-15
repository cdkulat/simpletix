package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"simpletix.kulat.co/internal/models"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file found")
	}
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	tickets  *models.TicketModel
}

func main() {
	connStr, exists := os.LookupEnv("CONNECTION_STRING")
	if !exists {
		log.Print("Connection string not found...")
	}

	addr := flag.String("addr", ":8080", "HTTP Network Address")      // Sets custom port when starting program - defaults to 8080
	dsn := flag.String("dsn", connStr, "postgres db connection pool") // sets a new flag for the database string
	flag.Parse()                                                      // scan cmd for flags

	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close() // for later, ensures that the database connection is closed before the main function exits

	// init instance of application struct
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		tickets:  &models.TicketModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(destination string) (*sql.DB, error) {
	db, err := sql.Open("postgres", destination)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
