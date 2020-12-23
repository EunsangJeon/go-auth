package oauth

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var githubOauthConfig = &oauth2.Config{
	Endpoint: github.Endpoint,
}

// Server runs OAuth authentication demonstration server
func Server() {
	err := godotenv.Load("oauth/server.env")
	if err != nil {
		log.Fatalf("Error loading .env file, %s", err)
	}

	githubOauthConfig.ClientID = os.Getenv("GITHUB_OAUTH_CLIENT_ID")
	githubOauthConfig.ClientSecret = os.Getenv("GITHUB_OAUTH_CLIENT_SECRET")

	http.HandleFunc("/", index)
	http.HandleFunc("/oauth/github", startGithubOauth)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>OAuth Example</title>
	</head>
	<body>
		<form action="/oauth/github" method="post">
			<input type="submit" value="Login with Github"/>
		</form>
	</body>
	</html>
`)
}

func startGithubOauth(w http.ResponseWriter, r *http.Request) {
	// You will need to use actual state to check. Usually you use DB to audit login states.
	redirectURL := githubOauthConfig.AuthCodeURL("test")
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}
