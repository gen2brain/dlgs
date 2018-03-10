// +build windows,!linux,!darwin,!js

package dlgs

import (
	"strings"
	"syscall"
	"unsafe"
)

// List displays a list dialog, returning the selected value and a bool for success.
func List(title, text string, items []string) (string, bool, error) {
	return listBox(title, text, "ClassList", items, false)
}

// ListMulti displays a multiple list dialog, returning the selected values and a bool for success.
func ListMulti(title, text string, items []string) ([]string, bool, error) {
	out, ok, err := listBox(title, text, "ClassListMulti", items, true)
	return strings.Split(out, lbSeparator), ok, err
}

// listBox displays list dialog.
func listBox(title, text, className string, items []string, multi bool) (string, bool, error) {
	var out string
	var hwndList syscall.Handle

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
				if multi {
					selCount := int(sendMessage(hwndList, lbGetSelCount, 0, 0))
					selItems := make([]int32, selCount)
					sendMessage(hwndList, lbGetSelItems, uintptr(selCount), uintptr(unsafe.Pointer(&selItems[0])))
					for i, idx := range selItems {
						out += items[idx]
						if i != len(selItems)-1 {
							out += lbSeparator
						}
					}
				} else {
					i := int(sendMessage(hwndList, lbGetCurSel, 0, 0))
					if i >= 0 {
						out = items[i]
					}
				}

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

	flags := wsBorder | wsChild | wsVisible | wsGroup | wsTabStop | esAutoVScroll
	if multi {
		flags |= lbsExtendedsel
	}
	hwndList, _ = createWindow(wsExClientEdge, "LISTBOX", title, uint64(flags), 10, 30, 200, 200, hwnd, 0, instance)

	hwndOK, _ := createWindow(wsExClientEdge, "BUTTON", "OK", wsChild|wsVisible|bsPushButton|wsGroup|wsTabStop, 10, 230, 90, 24, hwnd, 100, instance)
	hwndCancel, _ := createWindow(wsExClientEdge, "BUTTON", "Cancel", wsChild|wsVisible|bsPushButton|wsGroup|wsTabStop, 120, 230, 90, 24, hwnd, 110, instance)

	setWindowLong(hwnd, gwlStyle, getWindowLong(hwnd, gwlStyle)^wsMinimizeBox)
	setWindowLong(hwnd, gwlStyle, getWindowLong(hwnd, gwlStyle)^wsMaximizeBox)

	font := getMessageFont()
	sendMessage(hwndText, wmSetFont, font, 0)
	sendMessage(hwndList, wmSetFont, font, 0)
	sendMessage(hwndOK, wmSetFont, font, 0)
	sendMessage(hwndCancel, wmSetFont, font, 0)

	centerWindow(hwnd)

	for _, item := range items {
		i, _ := syscall.UTF16PtrFromString(item)
		sendMessage(hwndList, lbAddString, 0, uintptr(unsafe.Pointer(i)))
	}

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
