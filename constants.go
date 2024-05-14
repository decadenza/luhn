// Custom error definitions.

package luhn

import "errors"

// SECTION Configuration values.

// !SECTION

// SECTION Custom errors.
var ErrUnexpected = errors.New("unexpected")
var ErrInvalidBase = errors.New("invalid base")
var ErrInvalidValue = errors.New("invalid value")
var ErrInvalidChecksum = errors.New("invalid checksum")

// !SECTION
