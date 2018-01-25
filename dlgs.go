// Package dlgs is a cross-platform library for displaying dialogs and input boxes.
package dlgs

import (
	"errors"
	"runtime"
)

var verboseTests = false

var (
	// ErrUnsupported is returned when operating system is not supported.
	ErrUnsupported = errors.New("dlgs: unsupported operating system: " + runtime.GOOS)

	// ErrNotImplemented is returned when function is not implemented.
	ErrNotImplemented = errors.New("dlgs: function not implemented for " + runtime.GOOS)
)
