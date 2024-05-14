# Luhn algorithm
Implementation of Luhn algorithm for any base in the range 2 - 36.

## Notes for base 10 (most common)
The Luhn algorithm will detect all single-digit errors, as well as almost all transpositions of adjacent digits. It will not, however, detect transposition of the two-digit sequence 09 to 90 (or vice versa). It will detect most of the possible twin errors (it will not detect 22 ↔ 55, 33 ↔ 66 or 44 ↔ 77). 

## References
- [](https://en.wikipedia.org/wiki/Luhn_algorithm)
- [](https://en.wikipedia.org/wiki/Luhn_mod_N_algorithm)