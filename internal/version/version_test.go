package version

import "testing"

func TestGetVersion_WithDefaultVersion(t *testing.T) {
	expected := "TermOTP - Versão dev"

	if got := GetVersion(); got != expected {
		t.Errorf("Expected: '%s' but got: '%s'", expected, got)
	}
}

func TestGetVersion_WithCustomVersion(t *testing.T) {
	Version = "v0.2.0"
	expected := "TermOTP - Versão v0.2.0"

	if got := GetVersion(); got != expected {
		t.Errorf("Expected: '%s' but got: '%s'", expected, got)
	}
}
