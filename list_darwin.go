// +build darwin,!linux,!windows,!js

package dlgs

import (
	"os/exec"
	"strings"
	"syscall"
)

// List displays a list dialog, returning the selected value and a bool for success.
func List(title, text string, items []string) (string, bool, error) {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		return "", false, err
	}

	list := ""
	for i, l := range items {
		list += `"` + l + `"`
		if i != len(items)-1 {
			list += ", "
		}
	}

	o, err := exec.Command(osa, "-e", `choose from list {`+list+`} with prompt "`+text+`" with title "`+title+`"`).Output()
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
	osa, err := exec.LookPath("osascript")
	if err != nil {
		return []string{}, false, err
	}

	list := ""
	for i, l := range items {
		list += `"` + l + `"`
		if i != len(items)-1 {
			list += ", "
		}
	}

	o, err := exec.Command(osa, "-e", `choose from list {`+list+`} with multiple selections allowed with prompt "`+text+`" with title "`+title+`"`).Output()
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

	return strings.Split(out, ", "), ret, err
}
