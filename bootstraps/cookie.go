package bootstraps

import (
	"esaku-project/app/auths/models/web"
	"esaku-project/configs"
	"esaku-project/helpers"
	"esaku-project/types"
	"github.com/gorilla/securecookie"
	"net/http"
	"time"
)

type Cookie interface {
	getCookieToken() string
	getCookieRefresh() string
	getHashCookie() string
	getBlockCookie() string
	setSecureCookie() *securecookie.SecureCookie
	setTokenCookie(name string, data types.M, expiration time.Time, writer http.ResponseWriter)
	GetSecureCookie(name string, request *http.Request) (types.M, error)
	CreateTokenAndCookie(login web.LoginResponse, dataAccess, dataRefresh types.M, expirationAccess, expirationRefresh time.Time, writer http.ResponseWriter) error
	DeleteCookie(writer http.ResponseWriter) error
}

type CookieImpl struct {
	Config configs.Config
}

func NewCookieImpl(config configs.Config) *CookieImpl {
	return &CookieImpl{Config: config}
}

func (config *CookieImpl) getHashCookie() string {
	return config.Config.Get("COOKIE_HASH_KEY")
}

func (config *CookieImpl) getBlockCookie() string {
	return config.Config.Get("COOKIE_BLOCK_KEY")
}

func (config *CookieImpl) getCookieToken() string {
	return config.Config.Get("COOKIE_ACCESS_TOKEN")
}

func (config *CookieImpl) getCookieRefresh() string {
	return config.Config.Get("COOKIE_ACCESS_REFRESH_TOKEN")
}

func (config *CookieImpl) setSecureCookie() *securecookie.SecureCookie {
	secure := securecookie.New([]byte(config.getHashCookie()), []byte(config.getBlockCookie()))

	return secure
}

func (config *CookieImpl) setTokenCookie(name string, data types.M, expiration time.Time, writer http.ResponseWriter) {
	secure := config.setSecureCookie()

	encoded, err := secure.Encode(name, data)
	helpers.PanicIfError(err)

	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = encoded
	cookie.Expires = expiration
	cookie.Path = "/"
	cookie.HttpOnly = true

	http.SetCookie(writer, cookie)
}

func (config *CookieImpl) GetSecureCookie(name string, request *http.Request) (types.M, error) {
	secure := config.setSecureCookie()

	cookie, err := request.Cookie(name)

	helpers.PanicIfError(err)

	data := types.M{}

	if err = secure.Decode(name, cookie.Value, &data); err == nil {
		return data, nil
	}

	return nil, err
}

func (config *CookieImpl) CreateTokenAndCookie(login web.LoginResponse, dataAccess, dataRefresh types.M, expirationAccess, expirationRefresh time.Time, writer http.ResponseWriter) error {
	config.setTokenCookie(config.getCookieToken(), dataAccess, expirationAccess, writer)
	config.setTokenCookie(config.getCookieRefresh(), dataRefresh, expirationRefresh, writer)

	return nil
}

func (config *CookieImpl) DeleteCookie(writer http.ResponseWriter) error {
	config.setTokenCookie(config.getCookieToken(), types.M{}, time.Unix(0, 0), writer)
	config.setTokenCookie(config.getCookieRefresh(), types.M{}, time.Unix(0, 0), writer)

	return nil
}
