package main

import (
	"database/sql"
	"log"
	"log/slog"
	"net/http"
	"os"

	appConfig "github.com/Captain-Leftovers/beekeepers-log/config"
	"github.com/Captain-Leftovers/beekeepers-log/handler"
	"github.com/Captain-Leftovers/beekeepers-log/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
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

	TURSO_DATABASE_URL := os.Getenv("TURSO_DATABASE_URL")
	if TURSO_DATABASE_URL == "" {
		log.Fatal("$TURSO_DATABASE_URL must be set")
	}

	TURSO_AUTH_TOKEN := os.Getenv("TURSO_AUTH_TOKEN")
	if TURSO_AUTH_TOKEN == "" {
		log.Fatal("$TURSO_AUTH_TOKEN must be set")
	}

	TURSO_CONNECTION_STRING := os.Getenv("TURSO_CONNECTION_STRING")
	if TURSO_CONNECTION_STRING == "" {
		log.Fatal("$TURSO_CONNECTION_STRING must be set")
	}

	db, err := sql.Open("libsql", TURSO_CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}

	DBQ := database.New(db)

	appConfig := appConfig.NewAppConfig(DBQ)

	// TODO : implement different server structure capoable of passing dependencies to handlersa!!!!!!!!

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

	router.Get("/sign-up", handler.MakeHandler(handler.HandleSignUpIndex))
	router.Post("/sign-up", handler.MakeHandler(handler.HandlePOSTSignUpForm))

	router.Get("/sign-in", handler.MakeHandler(handler.HandleSignInIndex))
	router.Post("/sign-in", handler.MakeHandler(handler.HandleSignInProcess))

	slog.Info("Starting server on", "port", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, router))

}
