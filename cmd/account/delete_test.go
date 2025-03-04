package account

import (
	"os"
	"testing"

	"github.com/JVitoroliv3ira/termotp/internal/storage"
)

func TestDeleteAccount_WithValidData(t *testing.T) {
	testFile := "test_delete_valid.enc"
	storage.SetStorageFile(testFile)
	_ = os.Remove(testFile)

	err := AddAccount("MyService", "SECRET123", "Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	err = DeleteAccount("MyService", "Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	accounts, err := storage.LoadEncrypted("Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	_, getErr := accounts.GetAccount("MyService")
	if getErr == nil {
		t.Fatalf("Expected an error but got nil")
	}

	_ = os.Remove(testFile)
}

func TestDeleteAccount_WithInvalidName(t *testing.T) {
	err := DeleteAccount("", "Password123")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}
}

func TestDeleteAccount_WithInvalidPassword(t *testing.T) {
	err := DeleteAccount("MyService", "")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}
}

func TestDeleteAccount_WithNonExistingAccount(t *testing.T) {
	testFile := "test_delete_nonexisting.enc"
	storage.SetStorageFile(testFile)
	_ = os.Remove(testFile)

	err := DeleteAccount("NonExistingService", "Password123")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	_ = os.Remove(testFile)
}
