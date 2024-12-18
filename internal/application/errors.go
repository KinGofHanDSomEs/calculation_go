package application

import "errors"

var (
	ErrInvalidMethod  = errors.New("request method is specified incorrectly")
	ErrInternalServer = errors.New("internal server error")
)
