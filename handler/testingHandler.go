package handler

import (
	"net/http"

	testingView "github.com/Captain-Leftovers/beekeepers-log/view/testView"
)

// TODO : test routes below remove them if not needed

func TestColors() http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := testingView.Colors().Render(r.Context(), w)
		if err != nil {
			logError(r, err)
		}

	})

}
