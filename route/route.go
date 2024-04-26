package route

import (
	"net/http"

	"github.com/Captain-Leftovers/beekeepers-log/handler"
	"github.com/Captain-Leftovers/beekeepers-log/internal/database"
	"github.com/Captain-Leftovers/beekeepers-log/middleware"
	"github.com/go-chi/chi/v5"
)

func AddRoutes(
	mux *chi.Mux,
	JWT_SECRET string,
	DBQ *database.Queries,
) {

	router := mux

	router.Use(middleware.AddUserIfLoggedIn(JWT_SECRET, DBQ))

	router.Handle("/public/*", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	router.Get("/", handler.HandleHomeIndex())

	router.Get("/sign-in", handler.HandleSignInIndex())
	router.Post("/sign-in", handler.HandlePostSignIn(DBQ, JWT_SECRET))

	router.Get("/sign-up", handler.HandleSignUpIndex())
	router.Post("/sign-up", handler.HandlePostSignUpForm(DBQ, JWT_SECRET))

	router.Get("/sign-out", handler.LogOutHandler())
}
