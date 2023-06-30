package user

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sum28it/pass-manager/pkg/auth"
)

func Add(user User, secret string) ([]User, error) {

	if err := auth.Authenticate(secret, filepath.Join(Dir, localDir, envFile)); err != nil {
		return nil, err
	}

	var users []User
	users, err := read(os.O_RDWR)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, errors.New("error reading file")
		}
	}

	var matched []User
	for _, val := range users {
		if user.match(val) {
			matched = append(matched, val)
		}
	}

	if len(matched) != 0 {
		return matched, errors.New("User already exists")
	}

	// Encryption Here
	user.Password, err = encrypt(secret, user.Password)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error encrypting password")
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	user.ModifiedAt = now
	users = append(users, user)
	err = write(users)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
