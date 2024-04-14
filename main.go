package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("Initialization function")
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_CONNECTION  := os.Getenv("DB_CONNECTION")
	if DB_CONNECTION == "" {
		log.Fatal("DB_CONNECTION is not set in .env file")
	}	

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT is not set in .env file")
	}


}

func main() {
	// This is a comment
	println("Hello, World!")
}
