package routes

import (
	"net/http"

	"github.com/Captain-Leftovers/beekeepers-log/handler"
	"github.com/Captain-Leftovers/beekeepers-log/internal/database"
	"github.com/go-chi/chi/v5"
)

func AddRoutes(
	mux *chi.Mux,
	DBQ *database.Queries,
) {

	// TODO : see how do pass the DBQ to the handlers

	router := mux

	router.Handle("/public/*", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))

	router.Get("/sign-up", handler.MakeHandler(handler.HandleSignUpIndex))
	router.Post("/sign-up", handler.MakeHandler(handler.HandlePOSTSignUpForm))

	router.Get("/sign-in", handler.MakeHandler(handler.HandleSignInIndex))
	router.Post("/sign-in", handler.MakeHandler(handler.HandleSignInProcess))

}
