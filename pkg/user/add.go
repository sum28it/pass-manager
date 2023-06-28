package user

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sum28it/pass-manager/pkg/auth"
)

func Add(user User, secret string) error {

	if err := auth.Authenticate(secret, Dir+localDir+envFile); err != nil {
		return err
	}

	var users []User
	users, err := read(os.O_RDWR)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return errors.New("error reading file")
		}
	}

	// Encryption Here
	user.Password, err = encrypt(secret, user.Password)

	if err != nil {
		fmt.Println(err)
		return errors.New("error encrypting password")
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	user.ModifiedAt = now
	users = append(users, user)
	err = write(users)
	if err != nil {
		return err
	}
	return nil
}
