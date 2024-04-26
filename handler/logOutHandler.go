package handler

import "net/http"

func LogOutHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		r = HandleLogOutUser(w, r)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}
