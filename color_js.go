// +build js

package dlgs

import (
	"image/color"
	//"github.com/gopherjs/gopherjs/js"
)

// Color displays a color selection dialog, returning the selected color and a bool for success.
func Color(title, defaultColorHex string) (color.Color, bool, error) {
	return nil, false, ErrNotImplemented
}
