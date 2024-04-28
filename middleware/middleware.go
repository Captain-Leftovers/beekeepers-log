package middleware

import (
	"net/http"
	"strings"

	"github.com/Captain-Leftovers/beekeepers-log/internal/database"
	"github.com/Captain-Leftovers/beekeepers-log/util"
)

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

func AddUserIfLoggedIn(JWT_TOKEN string, DBQ *database.Queries) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if strings.Contains(r.URL.Path, "/public") {
				h.ServeHTTP(w, r)
				return
			}

			authUser, ok := util.GetUserFromAccessCookie(r, JWT_TOKEN)
			if !ok {
				authUser, ok = util.WriteNewAccessTokenToCookie(DBQ, w, r, JWT_TOKEN)
				if !ok {
					h.ServeHTTP(w, r)
					return
				}

			}

			r = util.AddUserToContext(r, authUser)

			h.ServeHTTP(w, r)
		})

	}
}

func RequireAuth() func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			user := util.GetUserFromContext(r.Context())
			if user.IsLoggedIn == true {

				next.ServeHTTP(w, r)
				return
			}

			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
		})
	}
}
