// Package luhn implements a simple manager for Luhn checksum calculation and validation.
package luhn

import (
	"strconv"
)

// Carries the chosen configuration and exposes the API.
type Manager struct {
	Base int
}

// Create a new checksum manager to create and validate coded strings.
// The base must be in the allowed range.
// It returns an instance that can be used to perform checksum calculation and validation.
func New(base int) (Manager, error) {

	err := isAllowed(base)
	if err != nil {
		return Manager{}, err
	}

	return Manager{Base: base}, nil
}

// Check that a code string is valid.
// The checksum must be the last character.
func (m Manager) IsValid(code string) bool {

	// Separate original checksum and payload.
	originalChecksum := string(code[len(code)-1])
	code = code[:len(code)-1]

	c, err := m.GetChecksum(code)
	if err != nil {
		return false
	}

	return c == originalChecksum
}

// Generate the checksum for a given input string.
// The input must *not* have a checksum suffix.
func (m Manager) GetChecksum(input string) (string, error) {

	var sum uint64

	if len(input) < 1 {
		// Empty string have no checksum.
		return "", ErrInvalidValue
	}

	base := uint64(m.Base)

	for i := int(0); i < len(input); i++ {
		// From the rightmost digit.
		// ParseUint will accept both lowercase and uppercase characters.
		// All characters are lowercase as used in strconv package.
		v, err := strconv.ParseUint(string(input[i]), m.Base, strconv.IntSize)
		if err != nil {
			return "", err
		}

		// Double the value of every second digit from the left.
		// For i = 1,3,5,etc.
		v *= (1 + uint64(i)%2)

		// Add digits together when above base. Only relevant when v >= base.
		// NOTE: Since we double each digit, the maximum possible value of v at this point is 2*(base-1).
		v = v/base + v%base

		sum += v

	}

	sum = (base - (sum % base)) % base // The smallest number (possibly zero) that must be added to sum to make it a multiple of the base.

	// FormatUint returns lowercase a-z characters. Enforcing uppercase (as more commonly used for codes).
	// TODO: Add an option in constructor.
	return strconv.FormatUint(sum, m.Base), nil
}

// Internal helper to validate possible bases.
// The Luhn mod N algorithm only works when N is divisible by 2.
func isAllowed(base int) error {

	if base < 2 || base > 36 {
		// Relying on the [strconv.FormatUint] function specs.
		// Range is 2 <= base <= 36.
		return ErrInvalidBase
	}

	if base%2 != 0 {
		// Base must be divisible by 2 for the Luhn algorithm to be effective.
		return ErrInvalidBase
	}

	return nil
}
