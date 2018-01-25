// +build linux,!windows,!darwin,!js

package dlgs

import (
	"os/exec"
	"strings"
	"syscall"
)

// List displays a list dialog, returning the selected value and a bool for success.
func List(title, text string, items []string) (string, bool, error) {
	cmd, err := cmdPath()
	if err != nil {
		return "", false, err
	}

	opts := []string{"--list", "--title", title, "--text", text, "--column=", "--hide-header"}
	args := append(opts, items...)

	o, err := exec.Command(cmd, args...).Output()
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

// ListMulti displays a multiple list dialog, returning the selected values and a bool for success.
func ListMulti(title, text string, items []string) ([]string, bool, error) {
	cmd, err := cmdPath()
	if err != nil {
		return []string{}, false, err
	}

	sep := "|"

	opts := []string{"--list", "--multiple", "--separator", sep, "--title", title, "--text", text, "--column=", "--hide-header"}
	args := append(opts, items...)

	o, err := exec.Command(cmd, args...).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return []string{}, ws.ExitStatus() == 0, nil
		}
	}

	ret := true
	out := strings.TrimSpace(string(o))
	if out == "" {
		ret = false
	}

	return strings.Split(out, sep), ret, err
}
