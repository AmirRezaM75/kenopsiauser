package kenopsiauser

import "errors"

var (
	UserNotVerified = errors.New("user not verified")
	UserNotFound    = errors.New("user not found")
)
