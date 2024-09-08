package errors

import (
	"encoding/json"
	"net/http"
)

// NewCustomError is used to create a new CustomError
func NewCustomError(statusCode int, msg string, details interface{}) CustomError {
	return CustomError{
		StatusCode: statusCode,
		Message:    msg,
		Details:    details,
	}
}

// ForbiddenError is mainly used for unauthorized access
// StatusCode is 403
func ForbiddenError(msg string) CustomError {
	return CustomError{
		StatusCode: http.StatusForbidden,
		Message:    msg,
	}
}

// InvalidRequestParsingError is used for invalid parsing of request body
// StatusCode is 400
func InvalidRequestParsingError(err error) CustomError {
	return CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid Request",
		Details:    err.Error(),
	}
}

// BadRequest is used to indicate that the request is invalid
// StatusCode is 400
func BadRequest(msg string) CustomError {
	return CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
	}
}

// ValidationError indicates that the request data is failed to validate
// StatusCode is 400
func ValidationError(msg string) CustomError {
	return CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
	}
}

// InternalError is used to indicate that an internal error occurred.
// Optionally, messages can be passed. Messages will be concatenated with a period if it's more than one.
// StatusCode is 500
func InternalError(err error, msgs ...string) CustomError {
	m := ""

	switch len(msgs) {
	case 0:
		m = "Internal Server Error"
	case 1:
		m = msgs[0]
	default:
		for _, v := range msgs {
			m += v + ". "
		}
	}

	return CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    m,
		Details:    err.Error(),
	}
}

// InternalServerError is used to indicate that an internal database error occurred
// StatusCode is 500
func InternalServerError(err error) CustomError {
	return CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Database Error",
		Details:    err.Error(),
	}
}

func EmptyRedisValue() CustomError {
	return CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    "Empty Redis Value",
	}
}

// Status returns the status code
func (e CustomError) Status() int {
	return e.StatusCode
}

// Error returns the error message
func (e CustomError) Error() string {
	return e.Message
}

func (e CustomError) Print() string {
	b, _ := json.Marshal(e)
	return string(b)
}
