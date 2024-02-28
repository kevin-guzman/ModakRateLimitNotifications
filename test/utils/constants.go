package utils

import (
	"errors"
)

var (
	ErrInfrastructure = errors.New("infrastructure error, like database connection error")
)
