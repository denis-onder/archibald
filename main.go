package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No environmental variable (.env) file found.")
	}
}

func main() {
	connectToDatabase()
	serve()
}
