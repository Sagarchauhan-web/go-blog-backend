package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/sagarchauhan-web/rssagg/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Starting...")

	godotenv.Load();

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the environment variable")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("dbURL is not found in the environment variable")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cant connect to database: ", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https//*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE","OPTIONS",},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter();

	v1Router.Get("/err", handlerErr)
	v1Router.Get("/users", apiCfg.handlerCreateUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	fmt.Println("Server Starting on Port:", portString)
	srv.ListenAndServe()
}