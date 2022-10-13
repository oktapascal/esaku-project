package middlewares

import (
	"esaku-project/exceptions"
	"net/http"
	"strings"
)

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		authHeader := strings.Split(request.Header.Get("Authorization"), "Bearer")

		if len(authHeader) != 2 {
			panic(exceptions.NewErrorUnauthorized("malformed jwt token"))
		}

		next.ServeHTTP(writer, request)
	})
}
