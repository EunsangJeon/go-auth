package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// UserClaims is custom claim to make jwt
type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

// Valid validtates if the claims are valid. It returns nil if valid.
func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("Token has expired")
	}

	if u.SessionID == 0 {
		return fmt.Errorf("Invalid session ID")
	}

	return nil
}

func createToken(c *UserClaims) (string, error) {
	key := []byte("anykey")

	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)

	signedToken, err := t.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("Error in createToken when signing token: %w", err)
	}

	return signedToken, nil
}

func main() {
	// TODO
}
