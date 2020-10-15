package errors

import (
	"net/http"
)

type ResError struct {
	Message string `json:"message"`
	Status int `json:"status"`
	Error string `json:"error"`
}

func NewBadRequestError(message string) *ResError{
	return &ResError{
		Message : message,
		Status : http.StatusBadRequest,
		Error: "bad_request",
	}
}

func NewNotFoundError (message string) *ResError{
	return &ResError{
		Message : message,
		Status : http.StatusNotFound,
		Error: "Data_not_found",
	}
}