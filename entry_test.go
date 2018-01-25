package dlgs

import (
	"testing"
)

func TestEntry(t *testing.T) {
	out, ret, err := Entry("Entry", "Enter something here, anything:", "default text")
	if err != nil {
		t.Error(err)
	}

	if verboseTests {
		println("out:", out, "ret:", ret)
	}
}
