package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Captain-Leftovers/beekeepers-log/internal/database"
	"github.com/Captain-Leftovers/beekeepers-log/types"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func AddUserToContext(r *http.Request, user types.AuthenticatedUser) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), types.UserKey, user))
}

func HashNewPassword(pass string) (string, error) {

	hashByte, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hashedPass := string(hashByte)

	return hashedPass, nil
}

func CompareHashToPass(hash string, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
}

type TokenType string

const (
	// TokenTypeAccess -
	TokenTypeAccess TokenType = "bee-access"
	// TokenTypeRefresh -
	TokenTypeRefresh TokenType = "bee-refresh"
)

// MakeJWT -
func MakeJWT(
	user types.JWTUser,
	tokenSecret string,
	expiresIn time.Duration,
	tokenType TokenType,
) (string, error) {
	signingKey := []byte(tokenSecret)

	if tokenType != TokenTypeAccess {
		user = types.JWTUser{
			ID: user.ID,
		}
	}

	customClaims := types.MyCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    string(tokenType),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
		},
		User: user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	return token.SignedString(signingKey)
}

// RefreshToken -
func RefreshToken(DBQ *database.Queries, ctx context.Context, tokenString string, tokenSecret string) (string, error) {
	claimsStruct := types.MyCustomClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) { return []byte(tokenSecret), nil },
	)
	if err != nil {
		return "", err
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return "", err
	}
	if issuer != string(TokenTypeRefresh) {
		return "", errors.New("invalid issuer")
	}

	dbUser, err := DBQ.GetUserById(ctx, claimsStruct.User.ID)
	if err != nil {
		return "", err
	}

	jwtUser := types.JWTUser{
		ID:       dbUser.ID,
		Email:    dbUser.Email,
		Username: dbUser.Username,
	}

	newToken, err := MakeJWT(
		jwtUser,
		tokenSecret,
		time.Hour*4,
		TokenTypeAccess,
	)
	if err != nil {
		return "", err
	}

	return newToken, nil
}

// ValidateJWT -
func ValidateJWTAndGetUser(tokenString, tokenSecret string) (types.JWTUser, error) {
	claimsStruct := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) { return []byte(tokenSecret), nil },
	)
	if err != nil {
		return types.JWTUser{}, err
	}

	if !token.Valid {
		return types.JWTUser{}, errors.New("invalid token")
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return types.JWTUser{}, err
	}
	if issuer != string(TokenTypeAccess) {
		return types.JWTUser{}, errors.New("invalid issuer")
	}

	user := types.JWTUser{}

	claims, ok := token.Claims.(*types.MyCustomClaims)

	if !ok {
		return types.JWTUser{}, errors.New("invalid claims")
	}

	user = claims.User

	return user, nil
}
