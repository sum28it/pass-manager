package user

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/sum28it/pass-manager/pkg/auth"
)

func Reset(secret string) error {

	// Authenticate user before resetting
	err := auth.Authenticate(secret, filepath.Join(Dir, localDir, envFile))
	if err != nil {
		return err
	}

	err = os.RemoveAll(filepath.Join(Dir, localDir))
	if err != nil {
		return errors.New("error deleting files")
	}
	return nil
}
