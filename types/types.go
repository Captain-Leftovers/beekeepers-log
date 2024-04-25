package types

import (
	"github.com/golang-jwt/jwt/v5"
)

type CookieName string

const (
	AccessCookie  CookieName = "bee-access"
	RefreshCookie CookieName = "bee-refresh"
)

type UserInCtx string

const UserKey UserInCtx = "user"

type AuthenticatedUser struct {
	ID         string
	Email      string
	Username   string
	IsLoggedIn bool
}

type JWTUser struct {
	ID       string
	Email    string
	Username string
}

type MyCustomClaims struct {
	jwt.RegisteredClaims
	User JWTUser `json:"user"`
}
