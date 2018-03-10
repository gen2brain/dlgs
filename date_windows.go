// +build windows,!linux,!darwin,!js

package dlgs

import (
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

// Date displays a calendar dialog, returning the date and a bool for success.
func Date(title, text string, defaultDate time.Time) (time.Time, bool, error) {
	out, ok, err := datePicker(title, text, defaultDate)
	if ok {
		tim, err := time.Parse("1/2/2006", out)
		if err != nil {
			return time.Now(), false, err
		}
		return tim, true, nil
	}

	return time.Now(), false, err
}

// datePicker displays calendar dialog.
func datePicker(title, text string, defaultDate time.Time) (string, bool, error) {
	var out string
	var hwndDate syscall.Handle

	className := "ClassDate"

	instance, err := getModuleHandle()
	if err != nil {
		return out, false, err
	}

	fn := func(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
		switch msg {
		case wmClose:
			destroyWindow(hwnd)
		case wmDestroy:
			postQuitMessage(0)
		case wmKeydown:
			if wparam == vkEscape {
				destroyWindow(hwnd)
			}
		case wmCommand:
			if wparam == 100 {
				var st systemtimeW
				sendMessage(hwndDate, dtmGetSystemTime, uintptr(gdtValid), uintptr(unsafe.Pointer(&st)))
				out = strconv.Itoa(int(st.wMonth)) + "/" + strconv.Itoa(int(st.wDay)) + "/" + strconv.Itoa(int(st.wYear))
				destroyWindow(hwnd)
			} else if wparam == 110 {
				destroyWindow(hwnd)
			}
		default:
			ret := defWindowProc(hwnd, msg, wparam, lparam)
			return ret
		}

		return 0
	}

	err = registerClass(className, instance, fn)
	if err != nil {
		return out, false, err
	}
	defer unregisterClass(className, instance)

	hwnd, _ := createWindow(0, className, title, wsOverlappedWindow, swUseDefault, swUseDefault, 235, 300, 0, 0, instance)
	hwndText, _ := createWindow(0, "STATIC", text, wsChild|wsVisible, 10, 10, 200, 16, hwnd, 0, instance)
	hwndDate, _ = createWindow(wsExClientEdge, "SysDateTimePick32", title, wsChild|wsVisible|wsGroup|wsTabStop, 10, 30, 200, 24, hwnd, 0, instance)
	hwndOK, _ := createWindow(wsExClientEdge, "BUTTON", "OK", wsChild|wsVisible|bsPushButton|wsGroup|wsTabStop, 10, 230, 90, 24, hwnd, 100, instance)
	hwndCancel, _ := createWindow(wsExClientEdge, "BUTTON", "Cancel", wsChild|wsVisible|bsPushButton|wsGroup|wsTabStop, 120, 230, 90, 24, hwnd, 110, instance)

	setWindowLong(hwnd, gwlStyle, getWindowLong(hwnd, gwlStyle)^wsMinimizeBox)
	setWindowLong(hwnd, gwlStyle, getWindowLong(hwnd, gwlStyle)^wsMaximizeBox)

	font := getMessageFont()
	sendMessage(hwndText, wmSetFont, font, 0)
	sendMessage(hwndDate, wmSetFont, font, 0)
	sendMessage(hwndOK, wmSetFont, font, 0)
	sendMessage(hwndCancel, wmSetFont, font, 0)

	st := systemtimeW{wYear: uint16(defaultDate.Year()), wMonth: uint16(defaultDate.Month()), wDay: uint16(defaultDate.Day()), wDayOfWeek: uint16(defaultDate.Weekday())}
	sendMessage(hwndDate, dtmSetSystemTime, uintptr(gdtValid), uintptr(unsafe.Pointer(&st)))

	centerWindow(hwnd)

	showWindow(hwnd, swShowNormal)
	updateWindow(hwnd)

	err = messageLoop(hwnd)
	if err != nil {
		return out, false, err
	}

	ret := false
	if out != "" {
		ret = true
	}

	return out, ret, nil
}
