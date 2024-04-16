package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Captain-Leftovers/beekeepers-log/handler"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func InitProg() error {
	return godotenv.Load()
   
}


func main() {

	if err := InitProg(); err != nil {
		log.Fatal(err)
	}
	
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("$PORT must be set")
	}





	router := chi.NewMux()


	router.Get("/", handler.HandleHomeIndex)

	slog.Info("Starting server on", "port", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, router))




}


