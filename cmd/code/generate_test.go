package code

import (
	"os"
	"testing"

	"github.com/JVitoroliv3ira/termotp/cmd/account"
	"github.com/JVitoroliv3ira/termotp/internal/storage"
)

func TestGenerateTOTP_WithValidData(t *testing.T) {
	testFile := "test_generate_valid.enc"
	storage.SetStorageFile(testFile)
	_ = os.Remove(testFile)

	err := account.AddAccount("MyService", "TESTSECRET", "Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	err = GenerateTOTP("MyService", "Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	_ = os.Remove(testFile)
}

func TestGenerateTOTP_WithInvalidName(t *testing.T) {
	err := GenerateTOTP("", "Password123")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}
}

func TestGenerateTOTP_WithInvalidPassword(t *testing.T) {
	err := GenerateTOTP("MyService", "")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}
}

func TestGenerateTOTP_WithNonExistingAccount(t *testing.T) {
	testFile := "test_generate_nonexisting.enc"
	storage.SetStorageFile(testFile)
	_ = os.Remove(testFile)

	err := GenerateTOTP("NoSuchService", "Password123")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	_ = os.Remove(testFile)
}

func TestGenerateTOTP_WithInvalidSecret(t *testing.T) {
	testFile := "test_generate_invalid_secret.enc"
	storage.SetStorageFile(testFile)
	_ = os.Remove(testFile)

	err := account.AddAccount("BrokenService", "ZZ!", "Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	err = GenerateTOTP("BrokenService", "Password123")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	_ = os.Remove(testFile)
}
