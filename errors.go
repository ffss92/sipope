package sipope

import "errors"

var (
	ErrUnauthorized = errors.New("invalid or expired api token")
	ErrNotFound     = errors.New("resource was not found")
)
