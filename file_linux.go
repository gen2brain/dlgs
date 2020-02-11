// +build linux,!windows,!darwin,!js

package dlgs

import (
	"os/exec"
	"strings"
	"syscall"
)

// File displays a file dialog, returning the selected file or directory, a bool for success, and an
// error if it was unable to display the dialog. Filter is a string that determines 
// which extensions should be displayed for the dialog. Separate multiple file 
// extensions by spaces and use "*.extension" format for cross-platform compatibility, e.g. "*.png *.jpg".
// A blank string for the filter will display all file types.
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

// FileMulti displays a file dialog that allows for selecting multiple files. It returns the selected 
// files, a bool for success, and an error if it was unable to display the dialog. Filter is a string 
// that determines which files should be available for selection in the dialog. Separate multiple file 
// extensions by spaces and use "*.extension" format for cross-platform compatibility, e.g. "*.png *.jpg".
// A blank string for the filter will display all file types.
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
