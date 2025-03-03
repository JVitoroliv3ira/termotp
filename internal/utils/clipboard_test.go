package utils

import (
	"errors"
	"testing"

	"github.com/atotto/clipboard"
)

func mockClipboardWriteSuccess() func() {
	original := clipboardWriteFunc
	clipboardWriteFunc = func(text string) error {
		return nil
	}
	return func() {
		clipboardWriteFunc = original
	}
}

func mockClipboardWriteFailure() func() {
	original := clipboardWriteFunc
	clipboardWriteFunc = func(text string) error {
		return errors.New("simulated clipboard error")
	}
	return func() {
		clipboardWriteFunc = original
	}
}

func mockClipboardUnsupported() func() {
	original := clipboard.Unsupported
	clipboard.Unsupported = true
	return func() {
		clipboard.Unsupported = original
	}
}

func TestCopyToClipboard_WithValidText(t *testing.T) {
	payload := "Hello, Clipboard!"

	restore := mockClipboardWriteSuccess()
	defer restore()

	err := CopyToClipboard(payload)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestCopyToClipboard_WithClipboardError(t *testing.T) {
	payload := "Hello, Clipboard!"

	restore := mockClipboardWriteFailure()
	defer restore()

	err := CopyToClipboard(payload)

	if err == nil {
		t.Errorf("Expected an error but got nil")
	}
}

func TestCopyToClipboard_WithUnsupportedClipboard(t *testing.T) {
	payload := "Hello, Clipboard!"
	expected := errors.New("não foi possível acessar a área de transferência")

	restore := mockClipboardUnsupported()
	defer restore()

	if got := CopyToClipboard(payload); got.Error() != expected.Error() {
		t.Errorf("Expected: '%v' but got: '%v'", expected, got)
	}
}
