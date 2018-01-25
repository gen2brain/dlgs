// +build windows,!linux,!darwin,!js

package dlgs

// Info displays information dialog.
func Info(title, text string) (bool, error) {
	ret := messageBox(title, text, mbOk|mbIconInfo)
	return ret == idOk, nil
}

// Warning displays warning dialog.
func Warning(title, text string) (bool, error) {
	ret := messageBox(title, text, mbOk|mbIconWarning)
	return ret == idOk, nil
}

// Error displays error dialog.
func Error(title, text string) (bool, error) {
	ret := messageBox(title, text, mbOk|mbIconError)
	return ret == idOk, nil
}

// Question displays question dialog.
func Question(title, text string, defaultCancel bool) (bool, error) {
	btn := mbDefaultIcon1
	if defaultCancel {
		btn = mbDefaultIcon2
	}

	ret := messageBox(title, text, mbYesNo|mbIconQuestion|btn)
	return ret == idYes, nil
}
