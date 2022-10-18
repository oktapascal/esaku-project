package middlewares

import (
	"context"
	"errors"
	"esaku-project/bootstraps"
	"esaku-project/exceptions"
	"esaku-project/types"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/securecookie"
	"net/http"
	"strings"
)

type MiddlewareAuth interface {
	MiddlewareCookie(next http.Handler) http.Handler
	MiddlewareRefreshToken(next http.Handler) http.Handler
	MiddlewareBearerToken(next http.Handler) http.Handler
}

type MiddlewareAuthImpl struct {
	CookieConfig bootstraps.Cookie
	JwtConfig    bootstraps.JWT
}

func NewMiddlewareAuthImpl(cookieConfig bootstraps.Cookie, jwtConfig bootstraps.JWT) *MiddlewareAuthImpl {
	return &MiddlewareAuthImpl{
		CookieConfig: cookieConfig,
		JwtConfig:    jwtConfig,
	}
}

func (middleware *MiddlewareAuthImpl) MiddlewareCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cookieAccess := middleware.CookieConfig.GetCookieToken()

		accessCookie, _ := request.Cookie(cookieAccess)

		if accessCookie == nil {
			next.ServeHTTP(writer, request)
		}

		dataAccess, err := middleware.CookieConfig.GetSecureCookie(cookieAccess, request)

		if err != nil && err != http.ErrNoCookie && err != securecookie.ErrMacInvalid {
			panic(errors.New(err.Error()))
		}

		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", dataAccess["value"]))

		next.ServeHTTP(writer, request)
	})
}

func (middleware *MiddlewareAuthImpl) MiddlewareRefreshToken(next http.Handler) http.Handler {
	//TODO implement me
	panic("implement me")
}

func (middleware *MiddlewareAuthImpl) MiddlewareBearerToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		authHeader := strings.Split(request.Header.Get("Authorization"), " ")
		if len(authHeader) != 2 {
			panic(exceptions.NewErrorUnauthorized("malformed jwt token"))
		}

		tokenJwt := authHeader[1]
		verifyJwt := middleware.JwtConfig.GetJWTKey()

		claims := &types.Claims{}

		tokenAccess, err := jwt.ParseWithClaims(tokenJwt, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(verifyJwt), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				panic(exceptions.NewErrorUnauthorized("signature token is invalid"))
				return
			}

			panic(exceptions.NewErrorUnauthorized(err.Error()))
			return
		}

		if !tokenAccess.Valid {
			panic(exceptions.NewErrorUnauthorized("token is invalid"))
		}

		data := types.M{"kode_lokasi": claims.KodeLokasi, "nik_input": claims.Nik}

		ctx := context.WithValue(request.Context(), "pic", data)

		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
