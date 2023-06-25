package auth

import "os"

func IsInit() bool {

	_, err := os.Stat("file/.env")
	return err == nil
}
