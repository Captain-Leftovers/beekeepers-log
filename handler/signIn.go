package handler

import (
	"time"

	"github.com/Captain-Leftovers/beekeepers-log/internal/database"
	"github.com/Captain-Leftovers/beekeepers-log/view/signIn"

	"net/http"
)

func HandleSignInIndex() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		r = HandleLogOutUser(w, r)

		err := signIn.SignInIndex().Render(r.Context(), w)

		if err != nil {
			logError(r, err)
		}
	})

}

func HandlePostSignIn(DBQ *database.Queries, JWT_SECRET string) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		time.Sleep(time.Second * 3)

		if err := r.ParseForm(); err != nil {
			logError(r, err)
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		err := HandleLogInUser(database.User{}, DBQ, JWT_SECRET, w, r, email, password)
		if err != nil {
			creds := signIn.SignInCredentials{
				Email: email,
			}

			errs := signIn.SignInErrors{
				Other: "Incorrect email or password",
			}

			err := signIn.SignInForm(creds, errs).Render(r.Context(), w)
			if err != nil {
				logError(r, err)
			}

			return

		}
	})
}
