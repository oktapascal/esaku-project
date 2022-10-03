package helpers

import "net/http"

func ReadFromRequestBodyMultipart(request *http.Request) {
	err := request.ParseMultipartForm(32 << 20) // max memory 32mb
	PanicIfError(err)
}
