package oauth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// In real world application, you will need DB of course
// Key is github ID, value is user ID
var githubConnections = make(map[string]string)

// JSON layout: {"data":{"viewer":{"id":"..."}}}
type githubResponse struct {
	Data struct {
		Viewer struct {
			ID string `json:"id"`
		} `json:"viewer"`
	} `json:"data"`
}

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
	http.HandleFunc("/oauth2/receive", completeGithubOauth)
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

func completeGithubOauth(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	state := r.FormValue("state")

	if state != "test" {
		http.Error(w, "State is incorrect", http.StatusBadRequest)
		return
	}

	token, err := githubOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Cloud not login", http.StatusInternalServerError)
		return
	}

	ts := githubOauthConfig.TokenSource(r.Context(), token)
	client := oauth2.NewClient(r.Context(), ts)

	// Github Oauth v4 uses GraphQL
	requestBody := strings.NewReader(`{"query":"query {viewer {id}}"}`)
	resp, err := client.Post("https://api.github.com/graphql", "application/json", requestBody)
	if err != nil {
		http.Error(w, "Could not get the user", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var gr githubResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		http.Error(w, "Github invalid response", http.StatusInternalServerError)
		return
	}

	githubID := gr.Data.Viewer.ID

	userID, ok := githubConnections[githubID]
	if !ok {
		// New User - create account
		// Maybe return, maybe not, depends
		// For now I just input any value
		userID = "testUser"
		githubConnections[githubID] = userID

	}
	fmt.Printf("%s: %s\n", githubID, userID)

	// Login to account userID using JWT
}
