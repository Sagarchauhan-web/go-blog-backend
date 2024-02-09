package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting...")

	godotenv.Load();

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the environment variable")
	}

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	fmt.Println("Server Starting on Port:", portString)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}