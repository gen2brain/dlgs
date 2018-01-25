// +build linux,!windows,!darwin,!js

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
	cmd, err := cmdPath()
	if err != nil {
		return nil, false, err
	}

	o, err := exec.Command(cmd, "--color-selection", "--title", title, "--color", defaultColorHex).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return nil, ws.ExitStatus() == 0, nil
		}
	}

	ret := true
	out := strings.TrimSpace(string(o))
	if out == "" {
		ret = false
	}

	return parseColor(out), ret, err
}

// parseColor returns color from output string.
func parseColor(out string) color.Color {
	col := color.RGBA{}

	if strings.HasPrefix(out, "#") {
		var r, g, b uint8
		fmt.Sscanf(out, "#%02x%02x%02x", &r, &g, &b)

		col.R = uint8(r)
		col.G = uint8(g)
		col.B = uint8(b)
	} else if strings.HasPrefix(out, "rgba(") {
		for _, s := range []string{"rgba", "(", ")"} {
			out = strings.Replace(out, s, "", -1)
		}
		t := strings.Split(out, ",")
		if len(t) == 4 {
			r, _ := strconv.ParseUint(t[0], 10, 8)
			g, _ := strconv.ParseUint(t[1], 10, 8)
			b, _ := strconv.ParseUint(t[2], 10, 8)
			a, _ := strconv.ParseUint(t[3], 10, 8)

			col.R = uint8(r)
			col.G = uint8(g)
			col.B = uint8(b)
			col.A = uint8(a)
		}
	} else if strings.HasPrefix(out, "rgb(") {
		for _, s := range []string{"rgb", "(", ")"} {
			out = strings.Replace(out, s, "", -1)
		}
		t := strings.Split(out, ",")
		if len(t) == 3 {
			r, _ := strconv.ParseUint(t[0], 10, 8)
			g, _ := strconv.ParseUint(t[1], 10, 8)
			b, _ := strconv.ParseUint(t[2], 10, 8)

			col.R = uint8(r)
			col.G = uint8(g)
			col.B = uint8(b)
		}
	}

	return col
}
