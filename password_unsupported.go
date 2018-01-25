// +build !linux,!windows,!darwin,!js

package dlgs

// Password displays a dialog, returning the entered value and a bool for success.
func Password(title, text string) (string, bool, error) {
	return "", false, ErrUnsupported
}
