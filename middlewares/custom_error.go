package middlewares

import (
	"errors"
	"esaku-project/exceptions"
	"net/http"
)

func CustomError(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				switch e := err.(type) {
				case string:
					err = errors.New(e)
				case error:
					err = e
				default:
					err = e
				}

				exceptions.ErrorHandler(writer, request, err)
			}
		}()

		next.ServeHTTP(writer, request)
	})
}
