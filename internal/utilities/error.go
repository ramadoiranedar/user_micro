package utilities

import (
	"net/http"

	"github.com/go-openapi/errors"
)

func SetError(code int, msg string, args ...interface{}) error {
	return errors.New(int32(code), msg)
}

func GetError(err error) errors.Error {
	if v, ok := err.(errors.Error); ok {
		return v
	}

	return errors.New(http.StatusInternalServerError, err.Error())
}
