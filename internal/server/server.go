package server

import (
	"net/http"

	"github.com/Captain-Leftovers/gohtmxtemplbeelog/internal/routes"
)

type Config struct {
	Host string
	Port int
}

func NewServer(
	config *Config,

) http.Handler {

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	return mux
}
