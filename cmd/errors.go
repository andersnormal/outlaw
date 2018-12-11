package cmd

import (
	"errors"
)

var (
	ErrNoProvider = errors.New("no domain provider enabled")
)
