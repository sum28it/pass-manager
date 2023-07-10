package auth

import "math/rand"

const (
	ALPHA int = iota
	NUM
	ALPHA_NUM
	ALPHA_NUM_SPECIAL
)

// Generates a random password of specified length.
// Argument code represents the type of password i.e. ALPHA, ALPHA_NUM, etc.
func GenerateRandomPassword(length, code int) string {
	password := ""
	var randomIndex int
	characters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789#$%&@")

	// To make sure that an alphabet and a digit is always present
	if code == ALPHA_NUM {
		password += string(characters[rand.Intn(52)]) + string(characters[52+rand.Intn(10)])
		length -= 2
	}
	// To make sure that an alphabet, a special character and a digit is always present
	if code == ALPHA_NUM_SPECIAL {
		password += string(characters[rand.Intn(52)]) + string(characters[52+rand.Intn(10)])
		password += string(characters[62+rand.Intn(5)])
		length -= 3
	}

	for i := 0; i < length; i++ {

		switch code {
		case ALPHA:
			randomIndex = rand.Intn(52)
		case NUM:
			randomIndex = 52 + rand.Intn(10)
		case ALPHA_NUM:
			randomIndex = rand.Intn(62)
		case ALPHA_NUM_SPECIAL:
			randomIndex = rand.Intn(len(characters))
		}
		password += string(characters[randomIndex])
	}
	return password
}
