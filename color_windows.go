// +build windows,!linux,!darwin,!js

package dlgs

import (
	"fmt"
	"image/color"
	"unsafe"
)

// Color displays a color selection dialog, returning the selected color and a bool for success.
func Color(title, defaultColorHex string) (color.Color, bool, error) {
	c, ok := colorDialog(defaultColorHex)
	if ok {
		col := color.RGBA{}
		col.R = byte(c & 0xff)
		col.G = byte((c >> 8) & 0xff)
		col.B = byte((c >> 16) & 0xff)
		return col, true, nil
	}

	return nil, false, nil
}

// colorDialog displays color dialog.
func colorDialog(defaultColorHex string) (uint32, bool) {
	var cc choosecolorW
	custom := make([]uint32, 16)

	cc.lpCustColors = &custom[0]
	cc.lStructSize = uint32(unsafe.Sizeof(cc))
	cc.flags = ccFullOpen | ccRgbInit

	var r, g, b uint8
	fmt.Sscanf(defaultColorHex, "#%02x%02x%02x", &r, &g, &b)
	cc.rgbResult = uint32(uint32(r) | uint32(g)<<8 | uint32(b)<<16)

	ok := chooseColor(&cc)
	if ok {
		return cc.rgbResult, true
	}

	return 0, false
}
