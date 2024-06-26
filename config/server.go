package config

import (
	"net/http"

	"github.com/Captain-Leftovers/beekeepers-log/internal/database"
	"github.com/Captain-Leftovers/beekeepers-log/route"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewServer(
	DBQ *database.Queries,
	JWT_SECRET string,

) http.Handler {
	mux := chi.NewMux()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not set
	})

	var handler = mux
	handler.Use(cors.Handler)

	route.AddRoutes(
		handler,
		JWT_SECRET,
		DBQ,
	)

	return handler
}
