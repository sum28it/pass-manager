package user

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/sum28it/pass-manager/pkg/auth"
)

func Init(secret string) (string, error) {

	// Check if app is already initialized
	if auth.IsInit(filepath.Join(Dir, localDir, envFile)) {
		return "", errors.New("app already initialized")
	}
	var err error

	// Directory holds the env file and passwords
	err = os.Mkdir(filepath.Join(Dir, localDir), 0644)
	if err != nil {
		return "", err
	}
	// Creating pass.json and .env files
	file, err := os.Create(filepath.Join(Dir, localDir, dataFile))
	if err != nil {
		return "", err
	}
	file.Close()
	// .env file holds the user secret and salt
	file, err = os.Create(filepath.Join(Dir, localDir, envFile))
	if err != nil {
		return "", err
	}

	// Creating a salt to append with passwords
	rand.New(rand.NewSource(time.Now().Unix()))
	salt := []byte{}
	for i := 0; i < 16; i++ {
		salt = append(salt, byte(rand.Int()%256))
	}

	h := sha256.New()
	h.Write([]byte(secret))

	// Write the salt and hashed secret
	file.WriteString(fmt.Sprintf("SALT=%x\n", string(salt)))
	file.WriteString(fmt.Sprintf("HASHED_SECRET=%x", string(h.Sum(nil))))
	file.Close()

	return filepath.Join(Dir, localDir), nil
}
