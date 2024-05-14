package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

type checksumManager struct {
	Base int
}

// Create a new checksum manager to create and validate coded strings.
func NewChecksumManager(base int) (checksumManager, error) {

	err := isAllowed(base)
	if err != nil {
		return checksumManager{}, err
	}

	return checksumManager{Base: base}, nil
}

// Check that a code string is valid.
// The checksum must be the last character.
func (m checksumManager) IsValid(code string) bool {

	originalChecksum := string(code[len(code)-1])
	payload := code[:len(code)-1]

	fmt.Println(code, "code")
	fmt.Println(payload, "payload")

	c, err := m.GetChecksum(payload) // Return uppercase checksum.
	if err != nil {
		return false
	}

	fmt.Println("original", originalChecksum, "new", c)

	return c == originalChecksum
}

// Generate the checksum for a given input string.
// The input must be *without* without checksum appended.
func (m checksumManager) GetChecksum(input string) (string, error) {

	var sum uint64
	inputLower := strings.ToLower(input)

	if len(input) < 1 {
		// Empty string have no checksum.
		return "", ErrInvalidValue
	}

	fmt.Println("input length", len(input))

	for i := uint64(len(input) - 1); i > 0; i-- {

		fmt.Println("index", i)

		// From the rightmost digit.
		// All characters are lowercase as used in strconv package.
		v, err := strconv.ParseUint(string(inputLower[i]), m.Base, strconv.IntSize)
		if err != nil {
			return "", err
		}

		// Double the value of every second digit from the left.
		// For i = 1,3,5,etc.
		v *= (1 + i%2)

		sum += v

	}

	base := uint64(m.Base)
	sum = base - (sum % base) // The smallest number (possibly zero) that must be added to sum to make it a multiple of the base.

	// FormatUint returns lowercase a-z characters. Enforcing uppercase (as more commonly used for codes).
	// TODO: Add an option in constructor.
	return strings.ToUpper(strconv.FormatUint(sum, m.Base)), nil
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
