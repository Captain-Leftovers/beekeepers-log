package main

import (
	"net/http"
)

func NewServer(
    // TODO : add logger and stores
// pointers to stors and loggers
) http.Handler {


    mux := http.NewServeMux()
    
    addRoutes(
        mux,
    )
    
    var handler http.Handler = mux

    // add middleware here for all routes to use

    handler = corsMiddleware(handler)


    return handler
}

