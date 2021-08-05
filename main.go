package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No environmental variable (.env) file found.")
	}
}

func main() {
	fmt.Println("Archibald started")
}
