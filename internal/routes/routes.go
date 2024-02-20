package routes

import (
	"net/http"

	"github.com/Captain-Leftovers/gohtmxtemplbeelog/internal/handler"
)

// pass data from newServer to registerRoutes if needed

func RegisterRoutes(mux *http.ServeMux) {
	// register fileserver from static/

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.Handle("/", handler.HomeHandler())

}
