package handler

import (
	"net/http"

	"github.com/Captain-Leftovers/beekeepers-log/view/profile"
)

func ProfileIndex() http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := profile.ProfileIndex().Render(r.Context(), w)
		if err != nil {
			logError(r, err)
		}

	})

}
