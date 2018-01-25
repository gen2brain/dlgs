// +build !linux,!windows,!darwin,!js

package dlgs

// Info displays information dialog box.
func Info(title, message string) (bool, error) {
	return false, ErrUnsupported
}

// Warning displays warning dialog box.
func Warning(title, message string) (bool, error) {
	return false, ErrUnsupported
}

// Error displays error dialog box.
func Error(title, message string) (bool, error) {
	return false, ErrUnsupported
}

// Question displays question dialog box.
func Question(title, text string, defaultCancel bool) (bool, error) {
	return false, ErrUnsupported
}
