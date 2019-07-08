// +build darwin,!linux,!windows,!js

package dlgs

import (
	"os/exec"
	"path"
	"strings"
	"syscall"
)

// File displays a file dialog, returning the selected file/directory and a bool for success.
func File(title, filter string, directory bool) (string, bool, error) {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		return "", false, err
	}

	f := "file"
	if directory {
		f = "folder"
	}

	t := ""
	if filter != "" {
		t = ` of type {"` + filter + `"}`
	}

	o, err := exec.Command(osa, "-e", `choose `+f+t+` with prompt "`+title+`"`).Output()
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

	tmp := strings.Split(out, ":")
	outPath := "/" + path.Join(tmp[1:]...)

	return outPath, ret, err
}

// FileMulti displays a file dialog, returning the selected files and a bool for success.
func FileMulti(title, filter string) ([]string, bool, error) {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		return []string{}, false, err
	}

	t := ""
	if filter != "" {
		t = ` of type {"` + filter + `"}`
	}

	o, err := exec.Command(osa, "-e", `choose file `+t+` with multiple selections allowed with prompt "`+title+`"`).Output()
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

	paths := make([]string, 0)

	l := strings.Split(out, ", ")
	for _, p := range l {
		tmp := strings.Split(p, ":")
		paths = append(paths, "/"+path.Join(tmp[1:]...))
	}

	return paths, ret, err
}
