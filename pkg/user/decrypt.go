package user

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/pbkdf2"
)

func decrypt(secret string, text string) (string, error) {

	err := godotenv.Load("files/.env")
	if err != nil {
		return "", errors.New("error loading env file")
	}

	salt := os.Getenv("SALT")

	// Derive an AES-256 key from the secret key using PBKDF2
	key := pbkdf2.Key([]byte(secret), []byte(salt), 4096, 32, sha256.New)

	// Create a new AES block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Create a new GCM (Galois/Counter Mode) instance
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	// Create a nonce
	nonce := make([]byte, gcm.NonceSize())

	decodedText, err := hex.DecodeString(text)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("error decrypting password")
	}
	// Decrypt the ciphertext
	decrypted, err := gcm.Open(nil, nonce, decodedText, nil)
	if err != nil {
		panic(err)
	}
	return string(decrypted), nil
}
