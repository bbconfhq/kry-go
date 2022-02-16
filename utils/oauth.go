package utils

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Oauth struct {
	UserApi string
	Cookie  CookieInfo
	Config  *oauth2.Config
}

type CookieInfo struct {
	Name       string
	Expiration time.Time
}

func GenerateOauthConfig(authType string) *Oauth {
	if authType == "Github" {
		return &Oauth{
			UserApi: "https://api.github.com/user",
			Cookie: CookieInfo{
				Name:       "oauth_state",
				Expiration: time.Now().Add(365 * 24 * time.Hour),
			},
			Config: &oauth2.Config{
				RedirectURL:  "http://localhost:8000/api/login/github/callback",
				ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
				ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
				Scopes:       []string{"read:user", "user:email"},
				Endpoint:     github.Endpoint,
			},
		}
	} else {
		return nil
	}
}

func (oauth Oauth) GenerateState(c echo.Context) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	cookie := http.Cookie{Name: oauth.Cookie.Name, Value: state, Expires: oauth.Cookie.Expiration}
	c.SetCookie(&cookie)

	return state
}

func (oauth Oauth) GetToken(code string) (*oauth2.Token, error) {
	return oauth.Config.Exchange(context.Background(), code)
}

func (oauth Oauth) GetUserBytes(token *oauth2.Token) ([]byte, error) {
	client := oauth.Config.Client(context.Background(), token)

	userInfoResp, err := client.Get(oauth.UserApi)
	if err != nil {
		return nil, err
	}
	defer userInfoResp.Body.Close()

	userInfo, err := ioutil.ReadAll(userInfoResp.Body)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
