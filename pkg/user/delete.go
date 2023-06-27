package user

import (
	"errors"
	"os"

	"github.com/sum28it/pass-manager/pkg/auth"
)

func Delete(user *User, secret string, force bool) error {

	// Authenticate user
	err := auth.Authenticate(secret, Dir+localDir+envFile)
	if err != nil {
		return err
	}

	// Filter users
	var users []User
	users, err = read(os.O_RDWR)
	if err != nil {
		return err
	}
	filtered := filter(user, users)
	if len(filtered) == len(users) {
		return errors.New("no such user exists")
	}

	if len(users)-len(filtered) > 1 && !force {
		return errors.New("more than one such user found")
	}
	err = write(filtered)
	if err != nil {
		return err
	}
	return nil

}

func filter(u *User, users []User) []User {
	var result []User

	for _, val := range users {
		if !u.match(val) {
			result = append(result, val)
		}
	}
	return result

}
