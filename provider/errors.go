package provider

import (
	"errors"
)

var (
	ErrInvalidID                   = errors.New("Invalid id")
	ErrInvalidDomainName           = errors.New("Invalid domain name")
	ErrInvalidDomainDate           = errors.New("Invalid domain date")
	ErrInvalidDomainRedirectTarget = errors.New("Invalid domain redirect target")
	ErrInvalidRedirectStatusCode   = errors.New("Invalid redirect http status code")
)
