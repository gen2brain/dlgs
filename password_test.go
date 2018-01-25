package dlgs

import (
	"testing"
)

func TestPassword(t *testing.T) {
	out, ret, err := Password("Password", "Enter your API key:")
	if err != nil {
		t.Error(err)
	}

	if verboseTests {
		println("out:", out, "ret:", ret)
	}
}
