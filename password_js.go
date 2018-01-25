// +build js

package dlgs

import (
	"github.com/gopherjs/gopherjs/js"
)

// Password displays a dialog, returning the entered value and a bool for success.
func Password(title, text string) (out string, ret bool, err error) {
	return Entry(title, text, "")
}
