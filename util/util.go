package util

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/Captain-Leftovers/beekeepers-log/internal/auth"
	"github.com/Captain-Leftovers/beekeepers-log/internal/database"
	"github.com/Captain-Leftovers/beekeepers-log/types"
)

func AddUserToContext(r *http.Request, user types.AuthenticatedUser) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), types.UserKey, user))
}

func GetUserFromContext(ctx context.Context) types.AuthenticatedUser {

	user, ok := ctx.Value(types.UserKey).(types.AuthenticatedUser)
	if !ok {
		return types.AuthenticatedUser{}
	}

	return user
}

func WriteNewAccessTokenToCookie(DBQ *database.Queries, w http.ResponseWriter, r *http.Request, JWT_TOKEN string) (types.AuthenticatedUser, bool) {
	cookie, err := r.Cookie(string(types.RefreshCookie))
	if err != nil {
		return types.AuthenticatedUser{}, false
	}

	token := cookie.Value

	if token == "" {
		return types.AuthenticatedUser{}, false
	}

	newAccessToken, authUser, err := auth.UseRefreshToken(DBQ, r.Context(), token, JWT_TOKEN)

	if err != nil {
		slog.Info("Error in WriteNewAccessTokenToCookie" + err.Error())
		return types.AuthenticatedUser{}, false
	}

	expIn := time.Hour * 4

	http.SetCookie(w, &http.Cookie{
		Name:     string(types.AccessCookie),
		Value:    newAccessToken,
		Path:     "/",
		Expires:  time.Now().Add(expIn),
		HttpOnly: true,
		// TODO : set secure to true when in production
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return authUser, true
}

func GetUserFromAccessCookie(r *http.Request, JWT_TOKEN string) (user types.AuthenticatedUser, ok bool) {

	cookie, err := r.Cookie(string(types.AccessCookie))
	if err != nil {
		return types.AuthenticatedUser{}, false
	}

	token := cookie.Value

	if token == "" {
		slog.Info("token in access cookie is empty")
		return types.AuthenticatedUser{}, false
	}

	JWTuser, err := auth.ValidateJWTAndGetUser(token, JWT_TOKEN)

	if err != nil {
		return types.AuthenticatedUser{}, false
	}

	user = types.AuthenticatedUser{
		ID:         JWTuser.ID,
		Email:      JWTuser.Email,
		Username:   JWTuser.Username,
		IsLoggedIn: true,
	}

	return user, true
}
