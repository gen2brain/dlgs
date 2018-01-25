// +build !linux,!windows,!darwin,!js

package dlgs

// File displays a file dialog, returning the selected file/directory and a bool for success.
func File(title, filter string, directory bool) (string, bool, error) {
	return "", false, ErrUnsupported
}

// FileMulti displays a file dialog, returning the selected files and a bool for success.
func FileMulti(title, filter string) ([]string, bool, error) {
	return []string{}, false, ErrUnsupported
}
