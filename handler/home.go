package handler

import (
	"net/http"

	"github.com/Captain-Leftovers/beekeepers-log/view/home"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) {

    home.Index().Render(r.Context(), w)
}