// +build darwin,!linux,!windows,!js

package dlgs

import (
	"os/exec"
	"strings"
	"syscall"
)

// Info displays information dialog.
func Info(title, text string) (bool, error) {
	return osaDialog(title, text, "note")
}

// Warning displays warning dialog.
func Warning(title, text string) (bool, error) {
	return osaDialog(title, text, "caution")
}

// Error displays error dialog.
func Error(title, text string) (bool, error) {
	return osaDialog(title, text, "stop")
}

// Question displays question dialog.
func Question(title, text string, defaultCancel bool) (bool, error) {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		return false, err
	}

	btn := "Yes"
	if defaultCancel {
		btn = "No"
	}

	out, err := exec.Command(osa, "-e", `set T to button returned of (display dialog "`+text+`" with title "`+title+`" buttons {"No", "Yes"} default button "`+btn+`")`).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return ws.ExitStatus() == 0, nil
		}
	}

	ret := false
	if strings.TrimSpace(string(out)) == "Yes" {
		ret = true
	}

	return ret, err
}

// osaDialog displays dialog.
func osaDialog(title, text, icon string) (bool, error) {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		return false, err
	}

	out, err := exec.Command(osa, "-e", `tell application "System Events" to display dialog "`+text+`" with title "`+title+`" buttons {"OK"} default button "OK" with icon `+icon+``).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return ws.ExitStatus() == 0, nil
		}
	}

	ret := false
	if strings.TrimSpace(string(out)) == "OK" {
		ret = true
	}

	return ret, err
}
