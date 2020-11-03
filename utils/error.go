package utils

import (
	"log"
	"net/http"
)

type ApiErr struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message error) *ApiErr {
	log.Println(message)
	return &ApiErr{
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewInternalServerError(message error) *ApiErr {
	log.Println(message)
	return &ApiErr{
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

func NewUnauthorizedError(message error) *ApiErr {
	log.Println(message)
	return &ApiErr{
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}