package service

import "errors"

var (
	ErrorInvalidURL        = errors.New("invalid URL")
	ErrorInvalidCode       = errors.New("invalid code")
	ErrorURLTooLong        = errors.New("URL is too long")
	ErrorEpicRareCollision = errors.New("all attempts to generate a unique code failed")
)
