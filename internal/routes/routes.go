package routes

import (
	"net/http"

	"github.com/Captain-Leftovers/gohtmxtemplbeelog/internal/handler"
)

// pass data from newServer to registerRoutes if needed

func RegisterRoutes(mux *http.ServeMux) {

	mux.Handle("/", handler.HomeHandler())

}
