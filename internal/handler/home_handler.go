package handler

import (
	"net/http"

	"github.com/Captain-Leftovers/gohtmxtemplbeelog/internal/view/homeview"
)

type HomeHandler struct{}

func (h *HomeHandler) HandleShowHome(w http.ResponseWriter, r *http.Request) {

	err := homeview.ShowHome().Render(r.Context(), w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
