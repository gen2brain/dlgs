// +build darwin,!linux,!windows,!js

package dlgs

import (
	"fmt"
	"image/color"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

// Color displays a color selection dialog, returning the selected color and a bool for success.
func Color(title, defaultColorHex string) (color.Color, bool, error) {
	var ur, ug, ub uint8
	fmt.Sscanf(defaultColorHex, "#%02x%02x%02x", &ur, &ug, &ub)

	r := strconv.Itoa(int(ur))
	g := strconv.Itoa(int(ug))
	b := strconv.Itoa(int(ub))

	o, err := osaExecute(`tell application "Finder"`, `activate`, `choose color default color {`+r+`, `+g+`, `+b+`}`, `end tell`)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return nil, ws.ExitStatus() == 0, nil
		}
	}

	out := strings.TrimSpace(o)
	return parseColor(out), true, err
}

// parseColor returns color from output string.
func parseColor(out string) color.Color {
	col := color.RGBA{}

	for _, s := range []string{"rgb", "(", ")"} {
		out = strings.Replace(out, s, "", -1)
	}
	t := strings.Split(out, ", ")
	if len(t) == 3 {
		r, _ := strconv.ParseUint(t[0], 10, 32)
		g, _ := strconv.ParseUint(t[1], 10, 32)
		b, _ := strconv.ParseUint(t[2], 10, 32)

		col.R = uint8(r)
		col.G = uint8(g)
		col.B = uint8(b)
	}

	return col
}
