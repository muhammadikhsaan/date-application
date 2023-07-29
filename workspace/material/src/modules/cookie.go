package modules

import (
	"net/http"
	"os"
	"time"
)

var (
	secure = os.Getenv("COOKIE_SECURE")
)

type cookie struct {
}

type Cookie interface {
	CookieLogin(name string, token string) *http.Cookie
	CookieIdentity(name string, value string) *http.Cookie
	RemoveCookie(name string) *http.Cookie
}

func NewCookie() Cookie {
	return &cookie{}
}

func (c *cookie) CookieLogin(name string, token string) *http.Cookie {
	client := new(http.Cookie)

	client.Name = name
	client.Value = token
	client.Path = "/"
	client.Secure = secure != "false"
	client.HttpOnly = true
	client.SameSite = http.SameSiteDefaultMode
	client.Expires = time.Now().AddDate(1, 0, 0)

	return client
}

func (c *cookie) CookieIdentity(name string, value string) *http.Cookie {
	client := new(http.Cookie)

	client.Name = name
	client.Value = value
	client.Path = "/"
	client.Secure = secure != "false"
	client.HttpOnly = true
	client.SameSite = http.SameSiteDefaultMode
	client.Expires = time.Now().AddDate(0, 6, 0)

	return client
}

func (c *cookie) RemoveCookie(name string) *http.Cookie {
	client := new(http.Cookie)

	client.Name = name
	client.Value = ""
	client.Path = "/"
	client.Expires = time.Now().AddDate(0, 0, -1)

	return client
}
