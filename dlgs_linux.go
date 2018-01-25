// +build linux,!windows,!darwin,!js

package dlgs

import (
	"errors"
	"os/exec"
)

// cmdPath looks for supported programs in PATH
func cmdPath() (string, error) {
	cmd, err := exec.LookPath("qarma")
	if err != nil {
		e := err
		cmd, err = exec.LookPath("zenity")
		if err != nil {
			return "", errors.New("dlgs: " + err.Error() + "; " + e.Error())
		}
	}

	return cmd, err
}
