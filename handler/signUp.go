package handler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/Captain-Leftovers/beekeepers-log/internal/auth"
	"github.com/Captain-Leftovers/beekeepers-log/internal/database"
	"github.com/Captain-Leftovers/beekeepers-log/view/signUp"
	"github.com/google/uuid"
)

func HandleSignUpIndex() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := signUp.SignUpIndex().Render(r.Context(), w)
		if err != nil {
			logError(r, err)
		}
	})

}

func HandlePostSignUpForm(DBQ *database.Queries) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			logError(r, err)
		}

		username := r.FormValue("username")

		email := r.FormValue("email")

		password := r.FormValue("password")

		confirmPassword := r.FormValue("confirm-password")

		formCreds := signUp.SignUpCredentials{
			Email:    email,
			Username: username,
		}
		formErr := signUp.SignUpErrors{}

		if password != confirmPassword {
			formErr.Password = "passwords don't match"
			err := signUp.SignUpForm(formCreds, formErr).Render(r.Context(), w)
			if err != nil {
				logError(r, err)
			}
		}

		// TODO : validate the form data

		hashedPassword, err := auth.HashNewPassword(password)

		if err != nil {
			formErr.Other = "Something went wrong. Please try again later."
			err := signUp.SignUpForm(formCreds, formErr).Render(r.Context(), w)
			if err != nil {
				logError(r, err)
			}

		}

		user, err := DBQ.CreateUser(r.Context(), database.CreateUserParams{
			ID:        uuid.New(),
			Username:  username,
			Email:     email,
			Password:  hashedPassword,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		})

		if err != nil {
			formErr.Other = err.Error()
			logError(r, err)
			err := signUp.SignUpForm(formCreds, formErr).Render(r.Context(), w)
			if err != nil {
				logError(r, err)
			}

		}

		slog.Info("User created", "user", user)

		w.Header().Set("HX-Redirect", "/")
	})
}
