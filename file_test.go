package dlgs

import (
	"runtime"
	"testing"
)

func TestFileDir(t *testing.T) {
	out, ret, err := File("Choose directory", "", true)
	if err != nil {
		t.Error(err)
	}

	if verboseTests {
		println("out:", out, "ret:", ret)
	}
}

func TestFile(t *testing.T) {
	var filter string
	switch runtime.GOOS {
	case "linux":
		filter = "Images (*.jpeg,*.png,*.gif) | *.jpg *.jpeg *.png *.gif"
	case "windows":
		// https://msdn.microsoft.com/en-us/library/windows/desktop/ms646839.aspx
		filter = "Images (*.jpeg,*.png,*.gif)\x00*.jpg;*.jpeg;*.png;*.gif\x00All Files (*.*)\x00*.*\x00\x00"
	case "darwin":
		// https://developer.apple.com/library/content/documentation/Miscellaneous/Reference/UTIRef/Articles/System-DeclaredUniformTypeIdentifiers.html
		filter = "public.image"
	}

	out, ret, err := File("Select file", filter, false)
	if err != nil {
		t.Error(err)
	}

	if verboseTests {
		println("out:", out, "ret:", ret)
	}
}

func TestFileMulti(t *testing.T) {
	out, ret, err := FileMulti("Select files", "")
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
