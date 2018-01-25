// +build !linux,!windows,!darwin,!js

package dlgs

// Date displays a calendar dialog, returning the date and a bool for success.
func Date(title, text string, defaultDate time.Time) (time.Time, bool, error) {
	return nil, false, ErrUnsupported
}
