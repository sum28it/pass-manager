package user

import (
	"errors"
	"os"

	"github.com/sum28it/pass-manager/pkg/auth"
)

func Reset(secret string) error {

	// Authenticate user before resetting
	err := auth.Authenticate(secret, Dir+localDir+envFile)
	if err != nil {
		return err
	}

	err = os.RemoveAll(Dir + localDir)
	if err != nil {
		return errors.New("error deleting files")
	}
	return nil
}
