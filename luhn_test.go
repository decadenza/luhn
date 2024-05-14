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

			// Generate a single digit error, and expect it to be detected.
			n, err := strconv.ParseUint(string(code[0]), b, strconv.IntSize)
			if err != nil {
				t.Fatal(err)
			}
			codeWrong := strconv.FormatUint((n+1)%uint64(b), b) + code[1:]

			// Check if the wrong code is actually wrong.
			valid = luhn.IsValid(codeWrong)
			fmt.Printf("%s is valid? %t\n", codeWrong, valid)
			if valid {
				t.Fatal(ErrUnexpected)
			}
		}
	}
}

func TestExample(t *testing.T) {

	base := 16
	myPayload := "14FAD"

	luhn, err := New(base)
	if err != nil {
		panic(err)
	}

	// Generate and concatenate checksum.
	checksum, err := luhn.GetChecksum(myPayload)
	if err != nil {
		t.Fatal(err)
	}
	fullCode := myPayload + checksum

	// Check its validity.
	valid := luhn.IsValid(fullCode)
	fmt.Printf("%s is valid? %t\n", fullCode, valid)

	// A mistyping will not be valid.
	wrongCode := "1AFAD" + checksum
	valid = luhn.IsValid(wrongCode)
	fmt.Printf("%s is valid? %t\n", wrongCode, valid)

}
