package Error

import "net/http"

type RestError struct {
	Message string  `json:"message"`
	Erro    string  `json:"error"`
	Code    int     `json:"code"`
	Causes  []Cause `json:"cause,omitempty"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestError) Error() string {
	return r.Message
}

func NewRestError(message string, err string, code int, causes []Cause) *RestError {
	return &RestError{
		Message: message,
		Erro:    err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Erro:    "BAD_REQUEST",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Cause) *RestError {
	return &RestError{
		Message: message,
		Erro:    "BAD_REQUEST",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Erro:    "INTERNAL_SERVER_ERROR",
		Code:    http.StatusInternalServerError,
	}
}

func NewForbiddenError(message string) *RestError {
	return &RestError{
		Message: message,
		Erro:    "FORBIDDEN",
		Code:    http.StatusForbidden,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Erro:    "NOT_FOUND",
		Code:    http.StatusNotFound,
	}
}
