package dlgs

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	out, ret, err := Date("Date", "Date of traveling:", time.Now())
	if err != nil {
		t.Error(err)
	}

	if verboseTests {
		println("out:", out.String(), "ret:", ret)
	}
}
