package errors

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	code         string
	msg          string
	httpStatus   int
	wrappedError error
}

func (e ApiError) Error() string {
	return fmt.Sprintf("[code: %s][msg: %s]", e.code, e.msg)
}

func (e ApiError) Code() string {
	return e.code
}

func (e *ApiError) Status() int {
	return e.httpStatus
}

func New(code, msg string, httpCode int, err error) ApiError {
	return ApiError{
		code:         code,
		msg:          msg,
		httpStatus:   httpCode,
		wrappedError: err,
	}
}

var InvalidHashLengthError = New("invalid_hash_length", "hash lenght must be greater than 0", http.StatusBadRequest, nil)
var InvalidInputError = New("invalid_input", "error validating input", http.StatusBadRequest, nil)
var JSONDecodingError = func(err error) error {
	return New("invalid_json", "error decoding json", http.StatusBadRequest, err)
}
var NotFoundError = func(err error) error {
	return New("not_found", "error not found", http.StatusNotFound, err)
}
