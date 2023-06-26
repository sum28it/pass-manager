package auth

import (
	"os"
)

func IsInit(filename string) bool {

	_, err := os.Stat(filename)
	return err == nil
}
