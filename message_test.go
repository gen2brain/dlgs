package dlgs

import (
	"testing"
)

func TestInfo(t *testing.T) {
	ret, err := Info("Info", "Lorem ipsum dolor sit amet.")
	if err != nil {
		t.Error(err)
	}

	if verboseTests {
		println("ret:", ret)
	}
}

func TestWarning(t *testing.T) {
	ret, err := Warning("Warning", "Incomplete information!")
	if err != nil {
		t.Error(err)
	}

	if verboseTests {
		println("ret:", ret)
	}
}

func TestError(t *testing.T) {
	ret, err := Error("Error", "Cannot divide by zero.")
	if err != nil {
		t.Error(err)
	}

	if verboseTests {
		println("ret:", ret)
	}
}

func TestQuestion(t *testing.T) {
	ret, err := Question("Question", "Are you sure you want to format this media?", true)
	if err != nil {
		t.Error(err)
	}

	if verboseTests {
		println("ret:", ret)
	}
}
