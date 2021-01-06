package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var (
	ctx       = context.Background()
	oauthConf = &oauth2.Config{
		ClientID:     "3afe434cf25b6d79f3af",
		ClientSecret: "781fcef110b3afb6c76a1d0b8a11bb896d24fe12",
		Scopes:       []string{"user:email", "repo"},
		Endpoint:     githuboauth.Endpoint,
		RedirectURL:  "http://localhost:4200/getToken",
	}
	// oauthConf oauth2.Config
	// random string for oauth2 API calls to protect against CSRF
	oauthStateString = "jayanthi-github-oAuth"
	code             string
)

const htmlIndex = `<html><body>
Please login with <a href="/login">GitHub</a> to proceed.
</body></html>
`

// display login page
func HandleMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(htmlIndex))
}

// check login status
func CheckLoginBeforeGetUser(response http.ResponseWriter, request *http.Request) {
	code = request.FormValue("code")
	_, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		oauthConf.RedirectURL = "http://" + request.Host + request.RequestURI
		http.Redirect(response, request, "/login", http.StatusTemporaryRedirect)
		return
	}
}

// /login
func HandleGitHubLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConf.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

///////////////////////////////////////////////////////////////////
/** Personal Access Token Approach */

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func ValidateToken(personalAccessToken string) bool {
	tokenSource := &TokenSource{
		AccessToken: personalAccessToken,
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		fmt.Printf("client.Users.Get() failed with '%s'\n", err)
		return false
	}
	d, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Printf("json.MarshlIndent() failed with %s\n", err)
		return false
	}
	fmt.Printf("User:\n%s\n", string(d))
	return true
}

type AuthResp struct {
	token string
	code  string
	state string
}

func GetAuthToken(response http.ResponseWriter, request *http.Request) {
	code := request.FormValue("code")
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		oauthConf.RedirectURL = "http://localhost:4200/getToken"
		http.Redirect(response, request, "/login", http.StatusTemporaryRedirect)
		return
	}

	jsonValue, _ := json.Marshal(token)
	response.WriteHeader(http.StatusOK)
	response.Write(jsonValue)
	return
}
