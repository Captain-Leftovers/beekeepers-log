package middleware

import "net/http"

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// set headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// allow preflight requests
		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}




func WithUser( h http.Handler ) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		
	})
}