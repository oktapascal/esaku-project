package helpers

import (
	"esaku-project/types"
	"github.com/gorilla/securecookie"
	"net/http"
	"os"
	"time"
)

type CookieConfigImpl struct {
}

func NewCookieConfigImpl() *CookieConfigImpl {
	return &CookieConfigImpl{}
}

func (cookie *CookieConfigImpl) Get(key string) string {
	return os.Getenv(key)
}

func (cookie *CookieConfigImpl) setSecureCookie() *securecookie.SecureCookie {
	hash := cookie.Get("COOKIE_HASH_KEY")
	block := cookie.Get("COOKIE_BLOCK_KEY")

	secure := securecookie.New([]byte(hash), []byte(block))

	return secure
}

func (cookie *CookieConfigImpl) SetCookieToken(writer http.ResponseWriter, cookieName string, data types.M, expiration time.Time) {
	secure := cookie.setSecureCookie()

	encoded, err := secure.Encode(cookieName, data)
	PanicIfError(err)

	newCookie := &http.Cookie{}

	newCookie.Name = cookieName
	newCookie.Value = encoded
	newCookie.Expires = expiration
	newCookie.HttpOnly = true
	newCookie.Secure = false
	newCookie.Path = "/"

	http.SetCookie(writer, newCookie)
}
