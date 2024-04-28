package handler

import (
	"net/http"

	"github.com/Captain-Leftovers/beekeepers-log/view/notFound"
)

func NotFoundHandler() http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notFound.NotFound().Render(r.Context(), w)
	})
}
