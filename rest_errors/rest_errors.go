package rest_errors

import (
	"errors"
	"net/http"
)

// RestErr the rest err interface for errors
type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

// NewError allows for specified/customized error reporting to the client
func NewError(msg string) error {
	return errors.New(msg)
}

// NewBadRequestError is for dynamic bad request errors
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NewNotFoundError is for dynamic bad request errors
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
		// Causes:  []interface{}{err},
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}
