// +build darwin,!linux,!windows,!js

package dlgs

import (
	"os/exec"
	"path"
	"strings"
	"syscall"
)

// File displays a file dialog, returning the selected file or directory, a bool for success, and an
// error if it was unable to display the dialog. Filter is a string that determines
// which extensions should be displayed for the dialog. Separate multiple file
// extensions by spaces and use "*.extension" format for cross-platform compatibility, e.g. "*.png *.jpg".
// A blank string for the filter will display all file types.
func File(title, filter string, directory bool) (string, bool, error) {
	f := "file"
	if directory {
		f = "folder"
	}

	t := ""
	if filter != "" {
		t = ` of type {`
		patterns := strings.Split(filter, " ")
		for i, p := range patterns {
			p = strings.Trim(p, "*.")
			t += osaEscapeString(p)
			if i < len(patterns)-1 {
				t += ", "
			}
		}
		t += "}"
	}

	o, err := osaExecute(`choose ` + f + t + ` with prompt ` + osaEscapeString(title))
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return "", ws.ExitStatus() == 0, nil
		}
	}

	out := strings.TrimSpace(o)
	tmp := strings.Split(out, ":")
	outPath := "/" + path.Join(tmp[1:]...)

	return outPath, true, err
}

// FileMulti displays a file dialog that allows for selecting multiple files. It returns the selected
// files, a bool for success, and an error if it was unable to display the dialog. Filter is a string
// that determines which files should be available for selection in the dialog. Separate multiple file
// extensions by spaces and use "*.extension" format for cross-platform compatibility, e.g. "*.png *.jpg".
// A blank string for the filter will display all file types.
func FileMulti(title, filter string) ([]string, bool, error) {
	t := ""
	if filter != "" {
		t = ` of type {`
		patterns := strings.Split(filter, " ")
		for i, p := range patterns {
			p = strings.Trim(p, "*.")
			t += osaEscapeString(p)
			if i < len(patterns)-1 {
				t += ", "
			}
		}
		t += "}"
	}
	o, err := osaExecute(`choose file ` + t + ` with multiple selections allowed with prompt ` + osaEscapeString(title))
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return []string{}, ws.ExitStatus() == 0, nil
		}
	}

	out := strings.TrimSpace(o)

	paths := make([]string, 0)

	l := strings.Split(out, ", ")
	for _, p := range l {
		tmp := strings.Split(p, ":")
		paths = append(paths, "/"+path.Join(tmp[1:]...))
	}

	return paths, true, err
}
