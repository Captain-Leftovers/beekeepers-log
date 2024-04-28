package handler

import (
	"net/http"

	"github.com/Captain-Leftovers/beekeepers-log/util"
	"github.com/Captain-Leftovers/beekeepers-log/view/profile"
)

func ProfileIndex() http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := util.GetUserFromContext(r.Context())

		err := profile.ProfileIndex(user).Render(r.Context(), w)
		if err != nil {
			logError(r, err)
		}

	})

}
