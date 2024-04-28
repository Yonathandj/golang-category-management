package exception

import (
	"golang-category-management/helper"
	"golang-category-management/model/response"
	"net/http"
)

func ExceptionHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	internalServerError(writer, request, err)
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	standardResponse := response.StandardResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Server Error",
		Data:       err,
	}
	helper.EncodeJSONBody(writer, standardResponse)
}
