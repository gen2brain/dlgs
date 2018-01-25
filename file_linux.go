// +build linux,!windows,!darwin,!js

package dlgs

import (
	"os/exec"
	"strings"
	"syscall"
)

// File displays a file dialog, returning the selected file/directory and a bool for success.
func File(title, filter string, directory bool) (string, bool, error) {
	cmd, err := cmdPath()
	if err != nil {
		return "", false, err
	}

	dir := ""
	if directory {
		dir = "--directory"
	}

	fileFilter := ""
	if filter != "" {
		fileFilter = "--file-filter=" + filter
	}

	o, err := exec.Command(cmd, "--file-selection", "--title", title, fileFilter, dir).Output()
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

// FileMulti displays a file dialog, returning the selected files and a bool for success.
func FileMulti(title, filter string) ([]string, bool, error) {
	cmd, err := cmdPath()
	if err != nil {
		return []string{}, false, err
	}

	sep := "|"

	fileFilter := ""
	if filter != "" {
		fileFilter = "--file-filter=" + filter
	}

	o, err := exec.Command(cmd, "--file-selection", "--multiple", "--separator", sep, "--title", title, fileFilter).Output()
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
