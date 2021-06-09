// +build darwin,!linux,!windows,!js

package dlgs

import (
	"os/exec"
	"strings"
	"syscall"
)

// Password displays a dialog, returning the entered value and a bool for success.
func Password(title, text string) (string, bool, error) {
	o, err := osaExecute(`set T to text returned of (display dialog ` + osaEscapeString(text) + ` with title ` + osaEscapeString(title) + ` default answer "" with hidden answer)`)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return "", ws.ExitStatus() == 0, nil
		}
	}

	out := strings.TrimSpace(o)

	return out, true, err
}
