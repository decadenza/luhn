package luhn

import (
	"fmt"
	"strconv"
)

func isValidMod16(code string) bool {

	originalC := string(code[len(code)-1])
	payload := code[:len(code)-1]

	c, err := generateChecksumMod16(payload)
	if err != nil {
		return false
	}

	return c == originalC
}

// The input without checksum.
func generateChecksumMod16(input string) (string, error) {

	var sum uint64

	maxIndex := uint64(len(input) - 1)

	for i := uint64(0); i <= maxIndex; i++ {

		// From the rightmost digit.
		v, err := strconv.ParseUint(string(input[maxIndex-i]), 16, 64)
		if err != nil {
			return "", err
		}

		// Double the value of every second digit.
		// For i = 1,3,5,etc.
		v *= 1 + i%2

		sum += v

	}

	sum = (16 - (sum % 16)) % 16

	return fmt.Sprintf("%X", sum), nil
}
