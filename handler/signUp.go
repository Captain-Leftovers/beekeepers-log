package handler

import (
	"log/slog"
	"net/http"

	"github.com/Captain-Leftovers/beekeepers-log/view/signUp"
)

func HandleSignUpIndex(w http.ResponseWriter, r *http.Request) error {
	return signUp.SignUpIndex().Render(r.Context(), w)
}

func HandlePOSTSignUpForm(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	username := r.FormValue("username")

	email := r.FormValue("email")

	password := r.FormValue("password")

	confirmPassword := r.FormValue("confirm-password")

	if password != confirmPassword {
		//TODO add error message
	}

	// TODO : validate the form data

	//print all the form data

	slog.Info("username: ", "username", username)
	slog.Info("email: ", "email", email)
	slog.Info("password: ", "password", password)
	slog.Info("confirmPassword: ", "confirmPassword", confirmPassword)

	return nil
}
