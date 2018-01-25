// +build js

package dlgs

import (
	"github.com/gopherjs/gopherjs/js"
)

// Entry displays input dialog, returning the entered value and a bool for success.
func Entry(title, text, defaultText string) (out string, ret bool, err error) {
	defer func() {
		e := recover()

		if e == nil {
			return
		}

		if e, ok := e.(*js.Error); ok {
			err = e
		} else {
			panic(e)
		}
	}()

	out = js.Global.Call("prompt", title+"\n\n"+text, defaultText).String()

	ret = true
	if out == "" {
		ret = false
	}

	return
}
