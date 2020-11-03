package utils

import "net/http"

type ApiErr struct {
	Message error `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message error) *ApiErr {
	return &ApiErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewInternalServerError(message error) *ApiErr {
	return &ApiErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

func NewUnauthorizedError(message error) *ApiErr {
	return &ApiErr{
		Message: message,
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}