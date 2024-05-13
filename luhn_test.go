package luhn

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBasicFunctions(t *testing.T) {

	// Generate a random HEX code.
	r := rand.New(rand.NewSource(0))
	code := fmt.Sprintf("%X", r.Int63())

	// Generate and concatenate checksum.
	checksum, err := generateChecksumMod16(code)
	if err != nil {
		t.Fatal(err)
	}
	code += checksum

	valid := isValidMod16(code)
	fmt.Printf("%s is valid? %t\n", code, valid)
	if !valid {
		t.Fatal(ErrUnexpected)
	}

	// Fake a mistyping, expect valid to be false.
	codeWrong := string(code[0]) + code
	valid = isValidMod16(codeWrong)
	fmt.Printf("%s is valid? %t\n", codeWrong, valid)
	if valid {
		t.Fatal(ErrUnexpected)
	}

}
