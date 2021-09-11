// +build js

package dlgs

// Password displays a dialog, returning the entered value and a bool for success.
func Password(title, text string) (out string, ret bool, err error) {
	return Entry(title, text, "")
}
