# Luhn algorithm
Implementation of Luhn Mod N algorithm for *any base (modulus)* divisible by 2 and in the range 2 - 36.

## Pros and cons
The Luhn algorithm will detect all single-digit errors, as well as almost all transpositions of adjacent digits. It will not, however, detect transposition of the two-digit sequence _base-1_-0 to 0-_base-1_ (or vice versa), e.g. swapping 09 with 90 in Mod 10. 
It will detect most of the possible twin errors, but not all of them (e.g. in Mod 10 it will not detect 22 ↔ 55, 33 ↔ 66 or 44 ↔ 77). 

## Install and import
Install with:
```
go get -u github.com/decadenza/luhn
```
Import as:
```
import "github.com/decadenza/luhn"
```

## Basic example
```
package main

import (
    "fmt"

    "github.com/decadenza/luhn"
)

func main() {
    base := 16
    myPayload := "14FAD"

    luhn, err := luhn.New(base)
    if err != nil {
        panic(err)
    }

    // Generate and concatenate checksum.
    checksum, err := luhn.GetChecksum(myPayload)
    if err != nil {
        panic(err)
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
```

Refer to test files for further examples.


## References
- [Luhn algorithm](https://en.wikipedia.org/wiki/Luhn_algorithm)
- [Luhn mod N algorithm](https://en.wikipedia.org/wiki/Luhn_mod_N_algorithm)
