package auth

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/Captain-Leftovers/beekeepers-log/internal/database"
	"github.com/Captain-Leftovers/beekeepers-log/types"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

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
func UseRefreshToken(DBQ *database.Queries, ctx context.Context, tokenString string, tokenSecret string) (string, types.AuthenticatedUser, error) {
	claimsStruct := types.MyCustomClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) { return []byte(tokenSecret), nil },
	)
	if err != nil {
		return "", types.AuthenticatedUser{}, err
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return "", types.AuthenticatedUser{}, err
	}
	if issuer != string(TokenTypeRefresh) {
		return "", types.AuthenticatedUser{}, errors.New("invalid issuer")
	}

	dbUser, err := DBQ.GetUserById(ctx, claimsStruct.User.ID)
	if err != nil {
		return "", types.AuthenticatedUser{}, err
	}

	jwtUser := types.JWTUser{
		ID:       dbUser.ID,
		Email:    dbUser.Email,
		Username: dbUser.Username,
	}

	newAccessToken, err := MakeJWT(
		jwtUser,
		tokenSecret,
		time.Hour*4,
		TokenTypeAccess,
	)
	if err != nil {
		return "", types.AuthenticatedUser{}, err
	}

	return newAccessToken, types.AuthenticatedUser{
		ID:         dbUser.ID,
		Email:      dbUser.Email,
		Username:   dbUser.Username,
		IsLoggedIn: true,
	}, nil
}

// ValidateJWT -
func ValidateJWTAndGetUser(tokenString, tokenSecret string) (types.JWTUser, error) {

	claimsStruct := types.MyCustomClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) { return []byte(tokenSecret), nil },
	)
	if err != nil {

		slog.Info("Error inside ValidateJWTAndGetUser" + err.Error())

		return types.JWTUser{}, err
	}

	if !token.Valid {

		return types.JWTUser{}, errors.New("invalid token")
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {

		slog.Info("Error inside ValidateJWTAndGetUser" + err.Error())
		return types.JWTUser{}, err
	}
	if issuer != string(TokenTypeAccess) {

		slog.Info("Error inside ValidateJWTAndGetUser", "invalid issuer", issuer)
		return types.JWTUser{}, errors.New("invalid issuer")
	}

	user := types.JWTUser{}

	user = claimsStruct.User

	return user, nil
}
