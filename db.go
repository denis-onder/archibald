package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type ConnectionInformation struct {
	username string
	password string
	host     string
	port     string
	dbName   string
}

var Database *sql.DB

func getConnectionInformation() (ConnectionInformation, error) {
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DATABASE")

	if username == "" {
		return ConnectionInformation{}, errors.New("Postgres username is required.")
	}

	if password == "" {
		return ConnectionInformation{}, errors.New("Postgres password is required.")
	}

	if host == "" {
		return ConnectionInformation{}, errors.New("Postgres host is required.")
	}

	if port == "" {
		return ConnectionInformation{}, errors.New("Postgres port is required.")
	}

	if dbName == "" {
		return ConnectionInformation{}, errors.New("Postgres database name is required.")
	}

	return ConnectionInformation{username, password, host, port, dbName}, nil
}

func buildConnectionString(c ConnectionInformation) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.username, c.password, c.host, c.port, c.dbName)
}

func connectToDatabase() {
	connInfo, err := getConnectionInformation()

	if err != nil {
		log.Printf("Postgres connection error:\n%s", err)
	}

	connString := buildConnectionString(connInfo)

	db, err := sql.Open("postgres", connString)

	if err != nil {
		log.Fatal(err)
	} else {
		Database = db
		fmt.Printf("Database connection established!\n%s\n", connString)
	}
}
