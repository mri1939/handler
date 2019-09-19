package handler

import (
	"net/http"
)

// message is a default message for the handler
type message struct {
	Content string `json:"msg"` // Monosodium Glutamate
}

// MessageWithCode handle the request with a message encoded in JSON
func MessageWithCode(statusCode int, msg string) http.Handler {
	return HandleJSON(statusCode, message{msg})
}

// NotImplemented returns a handler saying that the requested method in not implemented
func NotImplemented() http.Handler {
	return MessageWithCode(http.StatusNotImplemented, "Not Implemented")
}

// NotFound returns a handler saying that the requested URI was not found
func NotFound() http.Handler {
	return MessageWithCode(http.StatusNotFound, "Page Not Found")
}

// InternalServerError returns a handler saying that there was an error in our server
func InternalServerError() http.Handler {
	return MessageWithCode(http.StatusInternalServerError, "Sorry! There is a problem on our side.")
}

// NoData returns a handler saying that the requested data could not be found
func NoData() http.Handler {
	return MessageWithCode(http.StatusNotFound, "Sorry ! We couldn't find the data")
}

// Unauthorized returns a handler saying that the request was unauthorized
func Unauthorized() http.Handler {
	return MessageWithCode(http.StatusUnauthorized, "Unauthorized")
}

// BadRequest returns a handler saying that the request was malformed
func BadRequest() http.Handler {
	return MessageWithCode(http.StatusBadRequest, "You sent us a bad request.")
}
