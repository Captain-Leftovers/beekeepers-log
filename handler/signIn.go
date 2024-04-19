package handler

import (
	"github.com/Captain-Leftovers/beekeepers-log/view/signIn"

	"net/http"
)


func HandleSignInIndex(w http.ResponseWriter, r *http.Request) error {

	return signIn.SignInIndex().Render(r.Context(), w)
}
