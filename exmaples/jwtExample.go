package examples

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

var key2 = []byte("anykey")
var keys2 = map[string][]byte{}

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

	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)

	signedToken, err := t.SignedString(key2)
	if err != nil {
		return "", fmt.Errorf("Error in createToken when signing token: %w", err)
	}

	return signedToken, nil
}

func parseToken(signedToken string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(signedToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("Invalid signing algorithm")
		}

		// key rotation example
		// to rotate key, you will need actually key rotation system.
		// and when in createToken(), you need to use "current key".
		// when you craete new key you can use rand like:
		// newKey := make([]byte, 64)
		// _, err := io.ReadFull(rand.Reader, newKey)
		// and use uuid to key map. please refer to keys2 map[string][]byte

		// kid, ok := t.Header["kid"].(string)
		// if !ok {
		// 	return nil, fmt.Errorf("Invalid key ID")
		// }

		// k, ok := keys2[kid]
		// if !ok {
		// 	return nil, fmt.Errorf("Invalid key ID")
		// }

		// return k, nil

		// end of key rotation example

		return key2, nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error in parseToken while parsing token: %w", err)
	}

	if !t.Valid {
		return nil, fmt.Errorf("Error in parseToken, token is not valid")
	}

	return t.Claims.(*UserClaims), nil
}
