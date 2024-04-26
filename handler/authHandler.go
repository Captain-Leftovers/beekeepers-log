package handler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/Captain-Leftovers/beekeepers-log/internal/auth"
	"github.com/Captain-Leftovers/beekeepers-log/internal/database"
	"github.com/Captain-Leftovers/beekeepers-log/types"
	"github.com/Captain-Leftovers/beekeepers-log/util"
)

func HandleLogInUser(dbUser database.User, DBQ *database.Queries, JWT_SECRET string, w http.ResponseWriter, r *http.Request, email string, password string) error {

	if dbUser.ID == "" || dbUser.Email == "" || dbUser.Username == "" || dbUser.Password == "" {
		var err error
		dbUser, err = DBQ.GetUserByEmail(r.Context(), email)
		if err != nil {
			slog.Info("err from DbUser fetched from database in HandleLoginUser")
			return err
		}

		err = auth.CompareHashToPass(dbUser.Password, password)
		if err != nil {
			logError(r, err)
			return err
		}
	}

	jwtUser := types.JWTUser{
		ID:       dbUser.ID,
		Email:    dbUser.Email,
		Username: dbUser.Username,
	}

	expIn := time.Hour * 4

	token, err := auth.MakeJWT(jwtUser, JWT_SECRET, expIn, auth.TokenTypeAccess)

	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     string(types.AccessCookie),
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(expIn),
		HttpOnly: true,
		// TODO : set secure to true when in production
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	expIn = time.Hour * 24 * 7 * 4

	refreshToken, err := auth.MakeJWT(types.JWTUser{ID: dbUser.ID}, JWT_SECRET, expIn, auth.TokenTypeRefresh)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     string(types.RefreshCookie),
		Value:    refreshToken,
		Path:     "/",
		Expires:  time.Now().Add(expIn),
		HttpOnly: true,
		// TODO : set secure to true when in production
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusSeeOther)
	return nil
}

func HandleLogOutUser(w http.ResponseWriter, r *http.Request) *http.Request {
	http.SetCookie(w, &http.Cookie{
		Name:     string(types.AccessCookie),
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		// TODO : set secure to true when in production
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     string(types.RefreshCookie),
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		// TODO : set secure to true when in production
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	r = util.AddUserToContext(r, types.AuthenticatedUser{})

	return r
}
