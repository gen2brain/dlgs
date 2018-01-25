// +build linux,!windows,!darwin,!js

package dlgs

import (
	"os/exec"
	"strings"
	"syscall"
)

// Entry displays input dialog, returning the entered value and a bool for success.
func Entry(title, text, defaultText string) (string, bool, error) {
	cmd, err := cmdPath()
	if err != nil {
		return "", false, err
	}

	o, err := exec.Command(cmd, "--entry", "--title", title, "--text", text, "--entry-text", defaultText).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return "", ws.ExitStatus() == 0, nil
		}
	}

	ret := true
	out := strings.TrimSpace(string(o))
	if out == "" {
		ret = false
	}

	return out, ret, err
}
