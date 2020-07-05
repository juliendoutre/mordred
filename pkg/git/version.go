package git

import (
	"os/exec"

	"github.com/pkg/errors"
)

// IsInstalled checks the user Git install.
func IsInstalled() (bool, error) {
	if _, err := exec.Command("git", "version").Output(); err != nil {
		return false, errors.Wrap(err, "error checking git version")
	}

	return true, nil
}
