package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	Astatus  int    `json:"status"`
	Amessage string `json:"message"`
	Anerror  string `json:"error,omitempty"`
}

func (e *apiError) Status() int {
	return e.Astatus
}

func (e *apiError) Message() string {
	return e.Amessage
}
func (e *apiError) Error() string {
	return e.Anerror
}
func NewApiError(statusCode int, message string) ApiError {
	return &apiError{
		Astatus:  statusCode,
		Amessage: message,
	}
}

func NewApiErrorFromBytes(body []byte) (ApiError, error) {
	var result apiError
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.New("invalid json body")
	}
	return &result, nil
}

func NewNotFoundApiError(message string) ApiError {
	return &apiError{
		Astatus:  http.StatusNotFound,
		Amessage: message,
	}
}

func NewInternalServerError(message string) ApiError {
	return &apiError{
		Astatus:  http.StatusInternalServerError,
		Amessage: message,
	}
}

func NewBadRequestError(message string) ApiError {
	return &apiError{
		Astatus:  http.StatusBadRequest,
		Amessage: message,
	}
}
