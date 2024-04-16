package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Captain-Leftovers/beekeepers-log/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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

	// TODO : check cors later if problems with htmx or other front end stuff
	cors := cors.New(cors.Options{
        AllowedOrigins:   []string{"https://*", "http://*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300, // Maximum value not set
    })

    router.Use(cors.Handler)
	
	// static files
	router.Handle("/public/*", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))

	slog.Info("Starting server on", "port", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, router))




}


