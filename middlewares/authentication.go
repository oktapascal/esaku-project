package middlewares

import (
	"context"
	"errors"
	"esaku-project/app/auths/models/web"
	"esaku-project/bootstraps"
	"esaku-project/exceptions"
	"esaku-project/types"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/securecookie"
	"net/http"
	"strings"
	"time"
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
		accessCookie, err := request.Cookie(cookieAccess)

		if accessCookie == nil && err != nil {
			next.ServeHTTP(writer, request)
			return
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
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Context().Value("token") == nil {
			next.ServeHTTP(writer, request)
			return
		}

		values := request.Context().Value("token").(*jwt.Token)
		claims := values.Claims.(*types.Claims)

		if time.Unix(claims.ExpiresAt.Unix(), 0).Sub(time.Now()) < 15*time.Minute {
			cookieRefresh := middleware.CookieConfig.GetCookieRefresh()
			refreshCookie, err := request.Cookie(cookieRefresh)

			if err == nil && refreshCookie != nil {
				dataRefresh, err := middleware.CookieConfig.GetSecureCookie(cookieRefresh, request)
				if err != nil && err != http.ErrNoCookie && err != securecookie.ErrMacInvalid {
					panic(errors.New(err.Error()))
				}

				refreshToken := dataRefresh["value"].(string)
				verifyRefresh := middleware.JwtConfig.GetRefreshKey()

				tokenRefresh, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(verifyRefresh), nil
				})

				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						panic(exceptions.NewErrorUnauthorized("signature token is invalid"))
						return
					}

					panic(exceptions.NewErrorUnauthorized(err.Error()))
					return
				}

				if tokenRefresh != nil && tokenRefresh.Valid {
					loginResponse := web.LoginResponse{
						Nik:        claims.Nik,
						KodeLokasi: claims.KodeLokasi,
					}
					accessToken, expirationAccess, err := middleware.JwtConfig.GenerateAccessToken(loginResponse)
					if err != nil {
						panic(exceptions.NewErrorUnauthorized("token is incorrect"))
					}

					refreshToken, expirationRefresh, err := middleware.JwtConfig.GenerateRefreshToken(loginResponse)
					if err != nil {
						panic(exceptions.NewErrorUnauthorized("token is incorrect"))
					}

					dataAccess := types.M{"value": accessToken}
					dataRefresh := types.M{"value": refreshToken}

					_ = middleware.CookieConfig.CreateTokenAndCookie(loginResponse, dataAccess, dataRefresh, expirationAccess, expirationRefresh, writer)
				}
			}

		}

		next.ServeHTTP(writer, request)
	})
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

		ctx := context.WithValue(request.Context(), "token", tokenAccess)

		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
