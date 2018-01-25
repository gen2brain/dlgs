// +build windows,!linux,!darwin,!js

package dlgs

// Entry displays input dialog, returning the entered value and a bool for success.
func Entry(title, text, defaultText string) (string, bool, error) {
	return editBox(title, text, defaultText, "ClassEntry", false)
}
