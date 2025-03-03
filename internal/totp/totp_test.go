package totp

import (
	"errors"
	"testing"
	"time"
)

func mockGenerateTOTP(code string, err error) func() {
	original := generateTOTPFunc
	generateTOTPFunc = func(secret string, t time.Time) (string, error) {
		return code, err
	}
	return func() { generateTOTPFunc = original }
}

func TestGenerateTOTP_WithValidSecret(t *testing.T) {
	payload := "VALIDSECRET"
	expected := "123456"

	restore := mockGenerateTOTP(expected, nil)
	defer restore()

	got, _, err := GenerateTOTP(payload)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestGenerateTOTP_WithInvalidSecret(t *testing.T) {
	payload := "INVALIDSECRET"
	expected := errors.New("invalid TOTP secret")

	restore := mockGenerateTOTP("", expected)
	defer restore()

	_, _, got := GenerateTOTP(payload)

	if got.Error() != expected.Error() {
		t.Errorf("Expected '%v', got '%v'", expected, got)
	}
}
