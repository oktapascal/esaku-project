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

func (cookieConfig *CookieConfigImpl) Get(key string) string {
	return os.Getenv(key)
}

func (cookieConfig *CookieConfigImpl) setSecureCookie() *securecookie.SecureCookie {
	hash := cookieConfig.Get("COOKIE_HASH_KEY")
	block := cookieConfig.Get("COOKIE_BLOCK_KEY")

	secure := securecookie.New([]byte(hash), []byte(block))

	return secure
}

func (cookieConfig *CookieConfigImpl) SetCookieToken(writer http.ResponseWriter, cookieName string, data types.M, expiration time.Time) {
	secure := cookieConfig.setSecureCookie()

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

func (cookieConfig *CookieConfigImpl) GetCookieToken(cookieName string, request *http.Request) (types.M, error) {
	secure := cookieConfig.setSecureCookie()

	cookie, err := request.Cookie(cookieName)
	PanicIfError(err)

	data := types.M{}

	if err = secure.Decode(cookieName, cookie.Value, &data); err == nil {
		return data, nil
	}

	return nil, err
}
