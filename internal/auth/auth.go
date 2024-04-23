package auth

import "golang.org/x/crypto/bcrypt"

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
