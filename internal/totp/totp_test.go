package totp

import (
	"testing"
	"time"

	"github.com/pquerna/otp/totp"
)

func TestGenerateTOTP_WithValidSecret(t *testing.T) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Test",
		AccountName: "test@example.com",
	})
	if err != nil {
		t.Fatalf("Failed to generate test TOTP key: %v", err)
	}

	secret := key.Secret()
	got, _, err := GenerateTOTP(secret)

	if err != nil {
		t.Fatalf("GenerateTOTP() returned an unexpected error: %v", err)
	}
	if len(got) != 6 {
		t.Errorf("GenerateTOTP() = %q; want a 6-digit code", got)
	}
}

func TestGenerateTOTP_WithValidSecret_ReturnsValidRemainingTime(t *testing.T) {
	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "Test",
		AccountName: "test@example.com",
	})

	secret := key.Secret()
	_, got, err := GenerateTOTP(secret)

	if err != nil {
		t.Fatalf("GenerateTOTP() returned an unexpected error: %v", err)
	}
	if got < 0 || got > 30 {
		t.Errorf("GenerateTOTP() = %d; want between 0 and 30", got)
	}
}

func TestGenerateTOTP_WithInvalidSecret(t *testing.T) {
	secret := "invalid-key"
	got, remaining, err := GenerateTOTP(secret)

	if err == nil {
		t.Fatalf("GenerateTOTP() should have returned an error for an invalid secret, but got: %q, %d", got, remaining)
	}
}

func TestGenerateTOTP_WithSameSecret_GeneratesSameCodeInShortTime(t *testing.T) {
	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "Test",
		AccountName: "test@example.com",
	})

	secret := key.Secret()

	got1, _, err1 := GenerateTOTP(secret)
	time.Sleep(1 * time.Second)
	got2, _, err2 := GenerateTOTP(secret)

	if err1 != nil || err2 != nil {
		t.Fatalf("GenerateTOTP() returned an unexpected error: %v, %v", err1, err2)
	}
	if got1 != got2 {
		t.Errorf("GenerateTOTP() = %q, %q; want them to be the same", got1, got2)
	}
}
