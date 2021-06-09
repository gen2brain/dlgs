// +build darwin,!linux,!windows,!js

package dlgs

import (
	"os/exec"
	"strings"
)

// osaEscape escapes a string to be used in AppleScript
func osaEscapeString(unescaped string) string {
	escaped := strings.ReplaceAll(unescaped, "\\", "\\\\")
	escaped = strings.ReplaceAll(escaped, "\"", "\\\"")
	escaped = strings.ReplaceAll(escaped, "\n", "\\\n")
	return `"` + escaped + `"`
}

// osaExecute executes AppleScript
func osaExecute(command ...string) (string, error) {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		return "", err
	}

	out, err := exec.Command(osa, "-e", strings.Join(command, "\n")).Output()
	return string(out), err
}
