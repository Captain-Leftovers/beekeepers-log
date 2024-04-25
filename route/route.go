package route

import (
	"net/http"

	"github.com/Captain-Leftovers/beekeepers-log/handler"
	"github.com/Captain-Leftovers/beekeepers-log/internal/database"
	"github.com/go-chi/chi/v5"
)

func AddRoutes(
	mux *chi.Mux,
	JWT_SECRET string,
	DBQ *database.Queries,
) {

	// TODO : see how do pass the DBQ to the handlers

	router := mux

	router.Handle("/public/*", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	router.Get("/", handler.HandleHomeIndex())

	router.Get("/sign-in", handler.HandleSignInIndex())
	router.Post("/sign-in", handler.HandlePostSignIn(DBQ, JWT_SECRET))

	router.Get("/sign-up", handler.HandleSignUpIndex())
	router.Post("/sign-up", handler.HandlePostSignUpForm(DBQ))
}
