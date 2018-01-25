// +build js

package dlgs

import (
	"github.com/gopherjs/gopherjs/js"
)

// Info displays information dialog.
func Info(title, text string) (ret bool, err error) {
	return alertDialog(title, text, "\u24d8")
}

// Warning displays warning dialog.
func Warning(title, text string) (ret bool, err error) {
	return alertDialog(title, text, "\u26a0")
}

// Error displays error dialog.
func Error(title, text string) (ret bool, err error) {
	return alertDialog(title, text, "\u166e")
}

// Question displays question dialog.
func Question(title, text string, defaultCancel bool) (ret bool, err error) {
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

	ret = js.Global.Call("confirm", "\u2753 "+title+"\n\n"+text).Bool()

	return
}

// alert displays alert.
func alertDialog(title, text, level string) (ret bool, err error) {
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

	ret = true
	js.Global.Call("alert", level+" "+title+"\n\n"+text)

	return
}
