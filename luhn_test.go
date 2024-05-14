package luhn

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestAllAllowedModulos(t *testing.T) {

	// For each base divisible by 2 in the expected range.
	for b := 2; b <= 36; b += 2 {

		fmt.Printf("Testing base %d\n", b)

		luhn, err := New(b)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < 10; i++ {
			// Repeat tests a few times with pseudo-random values.

			// Generate a random integer and convert.
			input := strconv.FormatUint(rand.Uint64(), b)

			// Generate and concatenate checksum.
			checksum, err := luhn.GetChecksum(input)
			if err != nil {
				t.Fatal(err)
			}
			code := input + checksum

			valid := luhn.IsValid(code)
			fmt.Printf("%s is valid? %t\n", code, valid)
			if !valid {
				t.Fatal(ErrInvalidChecksum)
			}

			// TODO: Find an always-failing scheme.
			//
			// Fake a mistyping (just repeating the first character), expect it to be invalid.
			// codeWrong := string(code[0]) + code
			// valid = luhn.IsValid(codeWrong)
			// fmt.Printf("%s is valid? %t\n", codeWrong, valid)
			// if valid {
			// 	t.Fatal(ErrUnexpected)
			// }
		}
	}
}
