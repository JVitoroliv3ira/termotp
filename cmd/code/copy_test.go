package code

import (
	"os"
	"testing"

	"github.com/JVitoroliv3ira/termotp/cmd/account"
	"github.com/JVitoroliv3ira/termotp/internal/storage"
)

func TestCopyTOTP_WithValidData(t *testing.T) {
	testFile := "test_copytotp_valid.enc"
	storage.SetStorageFile(testFile)
	_ = os.Remove(testFile)

	err := account.AddAccount("MyService", "TESTSECRET", "Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	err = CopyTOTP("MyService", "Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	_ = os.Remove(testFile)
}

func TestCopyTOTP_WithInvalidName(t *testing.T) {
	err := CopyTOTP("", "Password123")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}
}

func TestCopyTOTP_WithInvalidPassword(t *testing.T) {
	err := CopyTOTP("MyService", "")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}
}

func TestCopyTOTP_WithNonExistingAccount(t *testing.T) {
	testFile := "test_copytotp_nonexisting.enc"
	storage.SetStorageFile(testFile)
	_ = os.Remove(testFile)

	err := CopyTOTP("NoSuchService", "Password123")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	_ = os.Remove(testFile)
}

func TestCopyTOTP_WithInvalidSecret(t *testing.T) {
	testFile := "test_copytotp_invalid_secret.enc"
	storage.SetStorageFile(testFile)
	_ = os.Remove(testFile)

	err := account.AddAccount("MyBrokenService", "Z!Z", "Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	err = CopyTOTP("MyBrokenService", "Password123")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	_ = os.Remove(testFile)
}
