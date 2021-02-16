// +build linux,!windows,!darwin,!js

package dlgs

import (
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// Date displays a calendar dialog, returning the date and a bool for success.
func Date(title, text string, defaultDate time.Time) (time.Time, bool, error) {
	cmd, err := cmdPath()
	if err != nil {
		return time.Now(), false, err
	}

	o, err := exec.Command(cmd, "--calendar", "--title", title, "--text", text,
		"--day", strconv.Itoa(defaultDate.Day()),
		"--month", strconv.Itoa(int(defaultDate.Month())),
		"--year", strconv.Itoa(defaultDate.Year()),
		`--date-format=%d/%m/%Y`).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return time.Now(), ws.ExitStatus() == 0, nil
		}
	}

	out := strings.TrimSpace(string(o))

	tim, err := time.Parse("02/01/2006", out)
	if err == nil {
		return tim, true, nil
	}

	return time.Now(), false, err
}
