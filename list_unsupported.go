// +build !linux,!windows,!darwin,!js

package dlgs

// List displays a list dialog, returning the selected value and a bool for success.
func List(title, text string, items []string) (string, bool, error) {
	return "", false, ErrUnsupported
}

// ListMulti displays a multiple list dialog, returning the selected values and a bool for success.
func ListMulti(title, text string, items []string) ([]string, bool, error) {
	return []string{}, false, ErrUnsupported
}
