package user

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/pbkdf2"
)

func encrypt(secret string, text string) (string, error) {

	// Generating encryption key from secret and salt

	err := godotenv.Load(filepath.Join(Dir, localDir, envFile))
	if err != nil {
		return "", errors.New("error loading env file")
	}

	salt := os.Getenv("SALT")

	// Derive an AES-256 key from the secret key using PBKDF2
	key := pbkdf2.Key([]byte(secret), []byte(salt), 4096, 32, sha256.New)

	// Create a new AES block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create a new GCM (Galois/Counter Mode) instance
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create a nonce
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return "", err
	}

	// Encrypt the plaintext
	ciphertext := gcm.Seal(nil, nonce, []byte(text), nil)

	// prepend the cipherTextwith nonce
	ciphertext = append(nonce, ciphertext...)

	// Return the hex encoded string for ciphertext
	return hex.EncodeToString(ciphertext), nil

}
