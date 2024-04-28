package helper

import (
	"encoding/json"
	"net/http"
)

func DecodeJSONBody(request *http.Request, result any) {
	err := json.NewDecoder(request.Body).Decode(&result)
	HelperPanic(err)
}

func EncodeJSONBody(writer http.ResponseWriter, result any) {
	writer.Header().Add("Content-Type", "application/json")

	err := json.NewEncoder(writer).Encode(result)
	HelperPanic(err)
}
