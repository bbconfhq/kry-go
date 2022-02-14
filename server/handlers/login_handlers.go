package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io/ioutil"
	"kry-go/server"
	"log"
	"net/http"
	"os"
	"time"
)

type LoginHandler struct {
	server *server.Server
}

const githubUserApi = "https://api.github.com/user"

var githubOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8000/login/github/callback",
	ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	Scopes:       []string{"read:user", "user:email"},
	Endpoint:     github.Endpoint,
}

func MakeLoginHandler(server *server.Server) *LoginHandler {
	return &LoginHandler{server: server}
}

func generateStateOauthCookie(c echo.Context) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauth", Value: state, Expires: expiration}
	c.SetCookie(&cookie)

	return state
}

func getUserDataFromGithub(code string) (string, error) {
	// Get request to a set URL
	req, err := http.NewRequest(
		"GET",
		githubUserApi,
		nil,
	)

	if err != nil {
		log.Panic("API Request creation failed")
	}

	// Set the Authorization header before sending the request
	// Authorization: token {header}
	authHeaderValue := fmt.Sprintf("token %s", code)
	req.Header.Set("Authorization", authHeaderValue)

	// Make the request
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Panic("User Information Request Failed")
	}

	// Read the response as a byte slice
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Convert byte slice to string and return
	return string(respBody), nil
}

func (loginHandler *LoginHandler) LoginGithub(c echo.Context) error {
	// Create oauthState cookie
	oauthState := generateStateOauthCookie(c)
	u := githubOauthConfig.AuthCodeURL(oauthState)
	return c.Redirect(302, u)
}

func (loginHandler *LoginHandler) LoginGithubCallback(c echo.Context) error {
	// Read oauthState from Cookie
	oauthState, _ := c.Cookie("oauth")

	if c.FormValue("state") != oauthState.Value {
		log.Println("Invalid OAuth GitHub State")
		return c.JSON(http.StatusInternalServerError, nil)
	}

	data, err := getUserDataFromGithub(c.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, data)
}
