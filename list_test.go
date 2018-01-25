package dlgs

import (
	"testing"
)

func TestList(t *testing.T) {
	out, ret, err := List("List", "Select item from list:", []string{"Bug", "New Feature", "Improvement"})
	if err != nil {
		t.Error(err)
	}

	if verboseTests {
		println("out:", out, "ret:", ret)
	}
}

func TestListMulti(t *testing.T) {
	out, ret, err := ListMulti("ListMulti", "Select languages from list:", []string{"PHP", "Go", "Python", "Bash"})
	if err != nil {
		t.Error(err)
	}

	if verboseTests {
		if len(out) > 0 {
			println("out length:", len(out), "out[0]:", out[0])
		}
		println("ret:", ret)
	}
}
