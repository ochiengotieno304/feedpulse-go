package utils

import "errors"

var (
	ErrorInvalidFeedID      = errors.New("invalid feed id provided")
	ErrorInternaServerError = errors.New("error sending your request")
)
