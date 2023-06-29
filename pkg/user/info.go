package user

import (
	"fmt"
	"path/filepath"

	"github.com/sum28it/pass-manager/pkg/auth"
)

// Returns info about the application data storage
func Info() string {

	if !auth.IsInit(filepath.Join(Dir, localDir, envFile)) {
		return fmt.Sprintf("The app is not initialized. After initialization, the data will be stored at %s in %s", Dir, localDir)
	}

	return filepath.Join(Dir, localDir)

}
