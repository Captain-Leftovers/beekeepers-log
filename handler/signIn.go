package handler

import (
	"github.com/Captain-Leftovers/beekeepers-log/view/signIn"

	"net/http"
)

func HandleSignInIndex(w http.ResponseWriter, r *http.Request) error {

	return signIn.SignInIndex().Render(r.Context(), w)
}

func HandleSignInProcess(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	email := r.FormValue("email")
	// password := r.FormValue("password")

	// TODO : validate the form data
	loginSuccess := false

	if loginSuccess {
		//TODO : redirect to home page
	} else {
		creds := signIn.SignInCredentials{
			Email: email,
		}

		errs := signIn.SignInErrors{
			Email:    "Email is not a valid email format or does not exist in the database",
			Password: "Password is incorrect format",
		}

		return signIn.SignInForm(creds, errs).Render(r.Context(), w)

	}
	return nil
}
