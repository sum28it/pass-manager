package auth

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Authenticate(secret string) error {

	if !IsInit() {
		return errors.New("app not initialized")
	}

	err := godotenv.Load("files/.env")
	if err != nil {
		return errors.New("error loading .env file")
	}
	hashedSecret := os.Getenv("HASHED_SECRET")
	h := sha256.New()
	h.Write([]byte(secret))
	genHashedSecret := fmt.Sprintf("%x", string(h.Sum(nil)))

	fmt.Printf("%s\n", hashedSecret)
	fmt.Printf("%s\n", genHashedSecret)
	if string(hashedSecret) != string(genHashedSecret) {
		return errors.New("incorrect secret")
	}

	return nil
}
