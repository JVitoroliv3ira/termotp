package version

import "testing"

func TestGetVersion_WithDefaultVersion(t *testing.T) {
	Version = "dev"
	want := "TermOTP - Versão dev"

	if got := GetVersion(); got != want {
		t.Errorf("GetVersion() = %q; want %q", got, want)
	}
}

func TestGetVersion_WithCustomVersion(t *testing.T) {
	Version = "v1.0.0"
	want := "TermOTP - Versão v1.0.0"

	if got := GetVersion(); got != want {
		t.Errorf("GetVersion() = %q; want %q", got, want)
	}
}
