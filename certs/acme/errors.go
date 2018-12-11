package acme

import (
	"errors"
)

var (
	errHostNotConfigured = errors.New("acme/autocert: host not configured")
)
