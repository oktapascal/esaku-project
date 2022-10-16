package helpers

import (
	"net/http"
	"time"
)

func SetCookieToken(writer http.ResponseWriter, cookieName string, cookieValue string, expiration time.Time) {
	cookie := &http.Cookie{}

	cookie.Name = cookieName
	cookie.Value = cookieValue
	cookie.Expires = expiration
	cookie.HttpOnly = true

	http.SetCookie(writer, cookie)
}
