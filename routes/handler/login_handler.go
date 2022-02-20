package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"kry-go/model/entity"
	"kry-go/server"
	"kry-go/service"
	"kry-go/utils"
	"log"
	"net/http"
)

type LoginHandler struct {
	server  *server.Server
	service *service.LoginService
}

func MakeLoginHandler(server *server.Server) *LoginHandler {
	return &LoginHandler{server: server}
}

func (loginHandler *LoginHandler) LoginUser(c echo.Context) error {
	session, _ := session.Get("session", c)
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	loginType := c.Param("loginType")
	if val, ok := session.Values[loginType]; ok {
		return c.JSONPretty(http.StatusOK, val, "  ")
	}

	return c.JSONPretty(http.StatusNotFound, nil, "  ")
}

func (loginHandler *LoginHandler) LoginOrRegisterUser(c echo.Context) error {
	session, _ := session.Get("session", c)
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	loginType := c.Param("loginType")
	oauth := utils.GenerateOauthConfig(loginType)

	// Read oauthState from Cookie
	oauthState, _ := c.Cookie(oauth.Cookie.Name)

	if c.FormValue("state") != oauthState.Value {
		log.Println("Invalid OAuth GitHub State")
		return c.JSONPretty(http.StatusInternalServerError, nil, "  ")
	}

	token, err := oauth.GetToken(c.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		return c.JSONPretty(http.StatusInternalServerError, err, "  ")
	}

	data, err := oauth.GetUserBytes(token)
	if err != nil {
		log.Println(err.Error())
		return c.JSONPretty(http.StatusInternalServerError, err, "  ")
	}

	// TODO: Unmarshal data to entity.User ...
	// TODO: Translate entity.User to proper type
	var userInfo entity.User
	json.Unmarshal(data, &userInfo)

	session.Values[loginType] = bytes.NewBuffer(data).String()
	session.Save(c.Request(), c.Response())

	return c.JSONPretty(http.StatusOK, userInfo, "  ")
}
