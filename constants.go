package luhn

import "errors"

// Custom errors.
var (
	ErrUnexpected      = errors.New("unexpected")
	ErrInvalidBase     = errors.New("invalid base")
	ErrInvalidValue    = errors.New("invalid value")
	ErrInvalidChecksum = errors.New("invalid checksum")
)
