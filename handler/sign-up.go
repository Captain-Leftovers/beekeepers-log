package handler

import (
	"net/http"

	"github.com/Captain-Leftovers/beekeepers-log/view/signUp"
)

func HandleSignUpIndex(w http.ResponseWriter, r *http.Request) error {
    return signUp.SignUpIndex().Render(r.Context(),w)
}