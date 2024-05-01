package main

import (
	"database/sql"
	"log"
	"log/slog"
	"net/http"
	"os"

	appConfig "github.com/Captain-Leftovers/beekeepers-log/config"
	"github.com/Captain-Leftovers/beekeepers-log/internal/database"
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

	JWT_SECRET := os.Getenv("JWT_SECRET")
	if JWT_SECRET == "" {
		log.Fatal("$JWT_SECRET must be set")
	}

	db, err := sql.Open("libsql", TURSO_CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}

	DBQ := database.New(db)

	srv := appConfig.NewServer(DBQ, JWT_SECRET)

	slog.Info("Starting server on", "port", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, srv))

}

// TODO : populate data in db after creating tables and migrations
