// +build windows,!linux,!darwin,!js

package dlgs

import (
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
)

// File displays a file dialog, returning the selected file or directory, a bool for success, and an
// error if it was unable to display the dialog. Filter is a string that determines
// which extensions should be displayed for the dialog. Separate multiple file
// extensions by spaces and use "*.extension" format for cross-platform compatibility, e.g. "*.png *.jpg".
// A blank string for the filter will display all file types.
func File(title, filter string, directory bool) (string, bool, error) {
	if directory {
		out, ok := dirDialog(title)
		return out, ok, nil
	}

	out, ok := fileDialog(title, filter, false)
	return out, ok, nil
}

// FileMulti displays a file dialog that allows for selecting multiple files. It returns the selected
// files, a bool for success, and an error if it was unable to display the dialog. Filter is a string
// that determines which files should be available for selection in the dialog. Separate multiple file
// extensions by spaces and use "*.extension" format for cross-platform compatibility, e.g. "*.png *.jpg".
// A blank string for the filter will display all file types.
func FileMulti(title, filter string) ([]string, bool, error) {
	out, ok := fileDialog(title, filter, true)

	files := make([]string, 0)

	if !ok {
		return files, ok, nil
	}

	l := strings.Split(out, "\x00")
	if len(l) > 1 {
		for _, p := range l[1:] {
			files = append(files, filepath.Join(l[0], p))
		}
	} else {
		files = append(files, out)
	}

	return files, ok, nil
}

// fileDialog displays file dialog.
func fileDialog(title, filter string, multi bool) (string, bool) {
	var ofn openfilenameW
	buf := make([]uint16, maxPath)

	t, _ := syscall.UTF16PtrFromString(title)

	ofn.lStructSize = uint32(unsafe.Sizeof(ofn))
	ofn.lpstrTitle = t
	ofn.lpstrFile = &buf[0]
	ofn.nMaxFile = uint32(len(buf))

	if filter != "" {
		ofn.lpstrFilter = utf16PtrFromString(filter)
	}

	flags := ofnExplorer | ofnFileMustExist | ofnHideReadOnly
	if multi {
		flags |= ofnAllowMultiSelect
	}
	ofn.flags = uint32(flags)

	if getOpenFileName(&ofn) {
		return stringFromUtf16Ptr(ofn.lpstrFile), true
	}

	return "", false
}

// dirDialog displays directory dialog.
func dirDialog(title string) (string, bool) {
	var bi browseinfoW
	buf := make([]uint16, maxPath)

	t, _ := syscall.UTF16PtrFromString(title)

	bi.title = t
	bi.flags = bifEditBox | bifNewDialogStyle

	lpItem := shBrowseForFolder(&bi)
	ok := shGetPathFromIDList(lpItem, &buf[0])
	if ok {
		return stringFromUtf16Ptr(&buf[0]), true
	}

	return "", false
}
