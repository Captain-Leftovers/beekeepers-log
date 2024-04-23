package handler

import (
	"net/http"

	"github.com/Captain-Leftovers/beekeepers-log/view/home"
)

func HandleHomeIndex() http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := home.Index().Render(r.Context(), w)
		if err != nil {
			logError(r, err)
		}

	})
}
