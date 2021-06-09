// +build darwin,!linux,!windows,!js

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
	o, err := osaExecute(`set defaultDate to do shell script "date -j -r `+strconv.Itoa(int(defaultDate.Unix()))+` +%m/%d/%Y"`,
		`set T to text returned of (display dialog `+osaEscapeString(text+` (mm/dd/yyyy)`)+` with title `+osaEscapeString(title)+` default answer defaultDate)`)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return time.Now(), ws.ExitStatus() == 0, nil
		}
	}

	ret := true
	out := strings.TrimSpace(o)
	if out == "" {
		ret = false
	}

	tim, err := time.Parse("01/02/2006", out)
	if err != nil {
		return time.Now(), false, err
	}

	return tim, ret, err
}
