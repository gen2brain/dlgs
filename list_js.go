// +build js

package dlgs

import (
//"github.com/gopherjs/gopherjs/js"
)

// List displays a list dialog, returning the selected value and a bool for success.
func List(title, text string, items []string) (string, bool, error) {
	return "", false, ErrNotImplemented
}

// ListMulti displays a multiple list dialog, returning the selected values and a bool for success.
func ListMulti(title, text string, items []string) ([]string, bool, error) {
	return []string{}, false, ErrNotImplemented
}
