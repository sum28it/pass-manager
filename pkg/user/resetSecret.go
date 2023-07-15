package user

import (
	"errors"
	"os"

	"github.com/sum28it/pass-manager/pkg/auth"
)

func ResetSecret(secret, newSecret string) error {

	// return if authentication fails
	if err := auth.Authenticate(secret, Dir+localDir+envFile); err != nil {
		return err
	}

	users, err := read(os.O_RDONLY)
	if err != nil {
		return errors.New("error reading user data")
	}

	// Store all passwords encrypted using new key and then update all the users
	// to avoid partial update due to some error while encrypring or decryping
	// any of the passwords
	newPasswords := make([]string, len(users))

	for i := range users {
		textPass, err := decrypt(secret, users[i].Password)
		if err != nil {
			return err
		}
		newPasswords[i], err = encrypt(newSecret, textPass)
		if err != nil {
			return err
		}
	}

	// Update the users with new passwords
	for i := range users {
		users[i].Password = newPasswords[i]
	}

	return nil
}
