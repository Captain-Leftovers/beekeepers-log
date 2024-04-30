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

	mainRouter := mux

	mainRouter.Use(middleware.AddUserIfLoggedIn(JWT_SECRET, DBQ))

	// Public Routes
	mainRouter.Group(func(r chi.Router) {

		mainRouter.Handle("/public/*", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

		mainRouter.NotFound(handler.NotFoundHandler())
		mainRouter.Get("/", handler.HandleHomeIndex())

		mainRouter.Get("/sign-in", handler.HandleSignInIndex())
		mainRouter.Post("/sign-in", handler.HandlePostSignIn(DBQ, JWT_SECRET))

		mainRouter.Get("/sign-up", handler.HandleSignUpIndex())
		mainRouter.Post("/sign-up", handler.HandlePostSignUpForm(DBQ, JWT_SECRET))

	})

	// Private Routes
	// Require Authentication
	mainRouter.Group(func(r chi.Router) {
		r.Use(middleware.RequireAuth())

		r.Get("/sign-out", handler.LogOutHandler())

		r.Get("/profile", handler.ProfileIndex())
		r.Post("/profile", handler.HandleProfilePost(DBQ))
		r.Get("/profile/change-password", handler.ChangePasswordFields())
	})

	//---------------------------------------------------
	// TODO : test routes below remove them if not needed

	mainRouter.Get("/colors", handler.TestColors())

	//---------------------------------------------------

}
