package utils

import (
	"errors"
	"os"
	"testing"

	"github.com/atotto/clipboard"
)

func TestCopyToClipboard_WithValidText(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		t.Skip("Ignorando teste de clipboard no GitHub Actions (ambiente sem suporte)")
	}

	text := "Texto de teste"
	if got := CopyToClipboard(text); got != nil {
		t.Errorf("CopyToClipboard() = %q; want nil", got)
	}
}

func TestCopyToClipboard_WhenClipboardUnsupported(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		t.Skip("Ignorando teste de clipboard no GitHub Actions (ambiente sem suporte)")
	}

	clipboard.Unsupported = true
	defer func() { clipboard.Unsupported = false }()

	text := "Texto de teste"
	got := CopyToClipboard(text)
	want := errors.New("não foi possível acessar a área de transferência")

	if got == nil || got.Error() != want.Error() {
		t.Errorf("CopyToClipboard() = %q; want %q", got, want)
	}
}
