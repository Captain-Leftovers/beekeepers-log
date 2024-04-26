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

		r = HandleLogOutUser(w, r)

		err := signUp.SignUpIndex().Render(r.Context(), w)
		if err != nil {
			logError(r, err)
		}

	})

}

func HandlePostSignUpForm(DBQ *database.Queries, JWT_SECRET string) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			logError(r, err)
		}
		// TODO : see validation and maybe validation packaage see how all user input is validated in go
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
			return
		}

		// TODO : validate the form data

		hashedPassword, err := auth.HashNewPassword(password)

		if err != nil {
			formErr.Other = "Something went wrong. Please try again later."
			err := signUp.SignUpForm(formCreds, formErr).Render(r.Context(), w)
			if err != nil {
				logError(r, err)
			}

			return

		}

		dbUser, err := DBQ.CreateUser(r.Context(), database.CreateUserParams{
			ID:        uuid.New().String(),
			Username:  username,
			Email:     email,
			Password:  hashedPassword,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		})

		if err != nil {
			// TODO : add proper message
			formErr.Other = "some constraint failed"
			logError(r, err)
			err := signUp.SignUpForm(formCreds, formErr).Render(r.Context(), w)
			if err != nil {
				logError(r, err)
			}

			slog.Info("User created", "user", dbUser)
			return

		}

		err = HandleLogInUser(dbUser, DBQ, JWT_SECRET, w, r, email, password)

		if err != nil {
			logError(r, err)

			formErr.Other = "Incorrect email or password"
			err := signUp.SignUpForm(formCreds, formErr).Render(r.Context(), w)
			if err != nil {
				logError(r, err)
			}

		}

	})
}
