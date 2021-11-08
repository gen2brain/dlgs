// +build darwin,!linux,!windows,!js

package dlgs

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

// List displays a list dialog, returning the selected value and a bool for success.
func List(title, text string, items []string) (string, bool, error) {
	list := ""
	for i, l := range items {
		if l == "false" {
			return "", false, fmt.Errorf("Cannot use 'false' in items, as it's reserved by osascript's returned value.")
		}
		list += osaEscapeString(l)
		if i != len(items)-1 {
			list += ", "
		}
	}

	o, err := osaExecute(`choose from list {` + list + `} with prompt ` + osaEscapeString(text) + ` with title ` + osaEscapeString(title))
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return "", ws.ExitStatus() == 0, nil
		}
	}

	out := strings.TrimSpace(o)
	if out == "false" {
		return "", false, nil
	}

	return out, true, err
}

// ListMulti displays a multiple list dialog, returning the selected values and a bool for success.
func ListMulti(title, text string, items []string) ([]string, bool, error) {
	list := ""
	for i, l := range items {
		if l == "false" {
			return nil, false, fmt.Errorf("Cannot use 'false' in items, as it's reserved by osascript's returned value.")
		}
		list += osaEscapeString(l)
		if i != len(items)-1 {
			list += ", "
		}
	}

	o, err := osaExecute(`choose from list {` + list + `} with multiple selections allowed with prompt ` + osaEscapeString(text) + ` with title ` + osaEscapeString(title))
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return []string{}, ws.ExitStatus() == 0, nil
		}
	}

	out := strings.TrimSpace(o)
	if out == "false" {
		return nil, false, nil
	}

	return strings.Split(out, ", "), true, err
}
