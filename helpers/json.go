package helpers

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBodyJson(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)

	PanicIfError(err)
}

func WriteToResponseBodyJson(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)

	PanicIfError(err)
}
