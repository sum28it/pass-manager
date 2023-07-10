package auth

import (
	"strings"
	"testing"
	"unicode"
)

func parseString(password, specialChar string) []bool {
	var alpha, num, special, other bool
	for _, char := range password {
		switch {
		case unicode.IsLetter(char):
			alpha = true
		case unicode.IsDigit(char):
			num = true
		case strings.ContainsRune(specialChar, char):
			special = true
		default:
			other = true
		}

	}
	return []bool{alpha, num, special, other}
}

func TestGen(t *testing.T) {
	var specialChar = "#$%&@"
	var alpha, num, special, other bool

	password := GenerateRandomPassword(10, ALPHA)
	bools := parseString(password, specialChar)
	_, num, special, other = bools[0], bools[1], bools[2], bools[3]
	if num || special || other {
		t.Errorf("Want alpha string of len 10, Got: %s", password)
	}

	password = GenerateRandomPassword(10, NUM)
	bools = parseString(password, specialChar)
	alpha, _, special, other = bools[0], bools[1], bools[2], bools[3]
	if alpha || special || other {
		t.Errorf("Want numeric string of len 10, Got: %s", password)
	}

	password = GenerateRandomPassword(10, ALPHA_NUM)
	bools = parseString(password, specialChar)
	_, _, special, other = bools[0], bools[1], bools[2], bools[3]
	if special || other {
		t.Errorf("Want alpha-numeric string of len 10, Got: %s", password)
	}

	password = GenerateRandomPassword(10, ALPHA_NUM_SPECIAL)
	bools = parseString(password, specialChar)
	_, _, _, other = bools[0], bools[1], bools[2], bools[3]
	if other {
		t.Errorf("Want alpha-numeric-special string of len 10, Got: %s", password)
	}

}
