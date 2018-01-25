// +build !linux,!windows,!darwin,!js

package dlgs

import (
	"image/color"
)

// Color displays a color selection dialog, returning the selected color and a bool for success.
func Color(title, defaultColorHex string) (color.Color, bool, error) {
	return nil, false, ErrUnsupported
}
