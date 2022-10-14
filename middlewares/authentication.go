package middlewares

import (
	"esaku-project/exceptions"
	"fmt"
	"net/http"
	"strings"
)

func MiddlewareCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		tokenCookie, _ := request.Cookie("access-token")

		if tokenCookie == nil {
			next.ServeHTTP(writer, request)
		}

		tokenValue := tokenCookie.Value

		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenValue))

		next.ServeHTTP(writer, request)
	})
}

func MiddlewareAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		authHeader := strings.Split(request.Header.Get("Authorization"), "Bearer")

		if len(authHeader) != 2 {
			panic(exceptions.NewErrorUnauthorized("malformed jwt token"))
		}

		next.ServeHTTP(writer, request)
	})
}
