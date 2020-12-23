package jwtauth

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type customClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

var key = []byte("keyishere")

// Server runs HMAC authentication demonstration server
func Server() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/submit", bar)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func getJWT(msg string) (string, error) {
	claims := customClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Second).Unix(),
		},
		Email: msg,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	signedString, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("Could not sign token: %w", err)
	}

	return signedString, nil
}

func bar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("email")
	if email == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	ss, err := getJWT(email)
	if err != nil {
		http.Error(w, "Could not getJWT", http.StatusInternalServerError)
		return
	}

	c := http.Cookie{
		Name:  "session",
		Value: ss,
	}

	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func foo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{}
	}

	ss := c.Value
	token, err := jwt.ParseWithClaims(
		ss,
		&customClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Cookie had not been signed with SHA256")
			}

			return key, nil
		},
	)

	isValid := err == nil && token.Valid

	message := "Not logged in"
	claims := &customClaims{}
	if isValid {
		message = "Logged in"
		claims = token.Claims.(*customClaims)
	}

	html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<meta http-equiv="X-UA-Compatible" content="ie=edge">
			<title>HMAC Example</title>
		</head>
		<body>
			<p>Cookie value: ` + c.Value + `</p>
			<p>` + message + `</p>
			<p>Session email: ` + claims.Email + `</p>
			<p>Session expires at: ` + fmt.Sprint(claims.ExpiresAt) + `</p>
			<form action="/submit" method="post">
				<input type="email" name="email" />
				<input type="submit" />
			</form>
		</body>
		</html>
	`
	io.WriteString(w, html)
}
