package exception

import (
	"golang-category-management/helper"
	"golang-category-management/model/response"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)
	
}

func internalServerError(writer http.ResponseWriter, _ *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	standardResponse := response.StandardResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Server Error",
		Data:       err,
	}
	helper.EncodeJSONBody(writer, standardResponse)
}

func notFoundError(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		standardResponse := response.StandardResponse{
			StatusCode: http.StatusNotFound,
			Message:    "Not Found",
			Data:       exception.Error,
		}
		helper.EncodeJSONBody(writer, standardResponse)
		return true
	} else {
		return false
	}
}