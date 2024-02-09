package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting...")

	godotenv.Load();

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the environment variable")
	}

	fmt.Println("Port:", portString)
}