package handler

import (
	"net/http"

	"github.com/Captain-Leftovers/beekeepers-log/view/hivemind"
)

func HandleHivemindIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := hivemind.HivemindIndex(hivemind.HivemindMain()).Render(r.Context(), w)

		if err != nil {
			logError(r, err)
		}
	}
}

func HandleHivemindHivemind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := hivemind.HivemindMain().Render(r.Context(), w)

		if err != nil {
			logError(r, err)
		}
	}
}

func HandleHivemindInspections() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := hivemind.InspectionsMain().Render(r.Context(), w)

		if err != nil {
			logError(r, err)
		}
	}
}

func HandleHivemindHives() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := hivemind.HivesMain().Render(r.Context(), w)

		if err != nil {
			logError(r, err)
		}
	}
}

func HandleHivemindFarms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := hivemind.FarmsMain().Render(r.Context(), w)

		if err != nil {
			logError(r, err)
		}
	}
}
