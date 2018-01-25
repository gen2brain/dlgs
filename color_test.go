package dlgs

import (
	"testing"
)

func TestColor(t *testing.T) {
	col, ret, err := Color("Pick color", "#BEBEBE")
	if err != nil {
		t.Error(err)
	}

	if verboseTests {
		if col != nil {
			r, g, b, a := col.RGBA()
			println("r:", r, "g:", g, "b:", b, "a:", a)
		}
		println("ret:", ret)
	}
}
