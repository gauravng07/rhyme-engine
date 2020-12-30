package httpext

import (
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	httpStatusCode int
	ErrorCode      string      `json:"error_code"`
	ErrorMessage   string      `json:"error_message"`
	Data           interface{} `json:"data,omitempty"`
	cause          error
}

func NewErrorResponse(statusCode int, errorCode string, errMsg string, cause error, data interface{}) ErrorResponse {
	return ErrorResponse{statusCode, errorCode, errMsg, data, cause}
}

func NewBadRequestError(errorCode string, message string, cause error, data interface{}) ErrorResponse {
	return NewErrorResponse(http.StatusBadRequest, errorCode, message, cause, data)
}

func (c ErrorResponse) Error() string {
	if c.cause != nil {
		return fmt.Sprintf("HttpStatusCode: %v, Error Code: %s, Error Message: %s, Cause: %+v", c.httpStatusCode, c.ErrorCode, c.ErrorMessage, c.cause)
	}
	return fmt.Sprintf("HttpStatusCode: %v, Error Code: %s, Error Message: %s", c.httpStatusCode, c.ErrorCode, c.ErrorMessage)
}

var InternalServerError = NewErrorResponse(http.StatusInternalServerError, UnknownInternalServiceErrorCode, "Failed to complete request due to some technical error", nil, nil)
