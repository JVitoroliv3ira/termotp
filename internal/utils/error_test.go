package utils

import (
	"errors"
	"testing"
)

func mockLogFatal() (restore func(), logCalled *bool) {
	original := logFatalFunc
	called := false

	logFatalFunc = func(v ...interface{}) {
		called = true
	}

	return func() { logFatalFunc = original }, &called
}

func TestHandleError_WithNilError(t *testing.T) {
	restore, logCalled := mockLogFatal()
	defer restore()

	HandleError(nil)

	if *logCalled {
		t.Errorf("Expected logFatalFunc NOT to be called for nil error")
	}
}

func TestHandleError_WithNonNilError(t *testing.T) {
	payload := errors.New("simulated error")

	restore, logCalled := mockLogFatal()
	defer restore()

	HandleError(payload)

	if !*logCalled {
		t.Errorf("Expected logFatalFunc to be called for error: '%v'", payload)
	}
}
