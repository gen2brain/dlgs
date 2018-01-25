// +build linux,!windows,!darwin,!js

package dlgs

import (
	"os/exec"
	"syscall"
)

// Info displays information dialog.
func Info(title, text string) (bool, error) {
	return cmdDialog(title, text, "info")
}

// Warning displays warning dialog.
func Warning(title, text string) (bool, error) {
	return cmdDialog(title, text, "warning")
}

// Error displays error dialog.
func Error(title, text string) (bool, error) {
	return cmdDialog(title, text, "error")
}

// Question displays question dialog.
func Question(title, text string, defaultCancel bool) (bool, error) {
	cmd, err := cmdPath()
	if err != nil {
		return false, err
	}

	dflt := ""
	if defaultCancel {
		dflt = "--default-cancel"
	}

	err = exec.Command(cmd, "--question", "--title", title, "--text", text, dflt).Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return ws.ExitStatus() == 0, nil
		}
	}

	return true, err
}

// cmdDialog displays dialog.
func cmdDialog(title, text, level string) (bool, error) {
	cmd, err := cmdPath()
	if err != nil {
		return false, err
	}

	err = exec.Command(cmd, "--"+level, "--title", title, "--text", text).Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return ws.ExitStatus() == 0, nil
		}
	}

	return true, err
}
