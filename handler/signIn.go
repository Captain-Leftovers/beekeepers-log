package handler

import (
	"github.com/Captain-Leftovers/beekeepers-log/view/signIn"

	"net/http"
)

func HandleSignInIndex() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := signIn.SignInIndex().Render(r.Context(), w)

		if err != nil {
			logError(r, err)
		}
	})

}

func HandlePostSignIn() http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			logError(r, err)
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

			err := signIn.SignInForm(creds, errs).Render(r.Context(), w)
			if err != nil {
				logError(r, err)
			}

		}

	})
}
