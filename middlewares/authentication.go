package middlewares

import (
	"context"
	"errors"
	"esaku-project/exceptions"
	"esaku-project/helpers"
	"esaku-project/types"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/securecookie"
	"net/http"
	"os"
	"strings"
)

func MiddlewareCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cookieConfig := helpers.NewCookieConfigImpl()

		cookieName := cookieConfig.Get("COOKIE_ACCESS_TOKEN")

		tokenCookie, _ := request.Cookie(cookieName)

		if tokenCookie == nil {
			next.ServeHTTP(writer, request)
		}

		data, err := cookieConfig.GetCookieToken(cookieName, request)

		if err != nil && err != http.ErrNoCookie && err != securecookie.ErrMacInvalid {
			panic(errors.New(err.Error()))
		}

		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", data["value"]))

		next.ServeHTTP(writer, request)
	})
}

func MiddlewareAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		verifyKey := []byte(os.Getenv("JWT_KEY_TOKEN"))

		authHeader := strings.Split(request.Header.Get("Authorization"), " ")
		if len(authHeader) != 2 {
			panic(exceptions.NewErrorUnauthorized("malformed jwt token"))
		}
		tokenBearer := authHeader[1]

		claims := &helpers.Claims{}
		token, err := jwt.ParseWithClaims(tokenBearer, claims, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				panic(exceptions.NewErrorUnauthorized("signature token is invalid"))
				return
			}

			panic(exceptions.NewErrorUnauthorized(err.Error()))
			return
		}

		if !token.Valid {
			panic(exceptions.NewErrorUnauthorized("token is invalid"))
		}

		data := types.M{"kode_lokasi": claims.KodeLokasi, "nik_input": claims.Nik}

		ctx := context.WithValue(request.Context(), "pic", data)

		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
