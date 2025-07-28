package handler

import (
	"errors"
)

var (
	// ErrInvalidId is returned when the ID is not a positive integer greater than zero
	ErrInvalidId = errors.New("invalid ID, must be a positive integer greater than zero")
	// ErrUnexpectedJSON is returned when the JSON is not valid
	ErrUnexpectedJSON = errors.New("unexpected JSON format, check the request body")
)
