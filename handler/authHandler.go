package handler

import (
	"net/http"
	"time"

	"github.com/Captain-Leftovers/beekeepers-log/internal/auth"
	"github.com/Captain-Leftovers/beekeepers-log/internal/database"
	"github.com/Captain-Leftovers/beekeepers-log/types"
)

func HandleLogInUser(DBQ *database.Queries, JWT_SECRET string, w http.ResponseWriter, r *http.Request, email string, password string) error {

	dbUser, err := DBQ.GetUserByEmail(r.Context(), email)

	if err != nil {
		return err
	}

	err = auth.CompareHashToPass(dbUser.Password, password)
	if err != nil {
		logError(r, err)
		return err
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
