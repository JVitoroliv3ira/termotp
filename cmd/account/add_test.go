package account

import (
	"os"
	"testing"
	"time"

	"github.com/JVitoroliv3ira/termotp/internal/storage"
)

func TestAddAccount_WithValidData(t *testing.T) {
	testFile := "test_valid.enc"
	storage.SetStorageFile(testFile)
	_ = os.Remove(testFile)

	err := AddAccount("MyService", "SECRET123", "Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	accounts, err := storage.LoadEncrypted("Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	foundAccount, err := accounts.GetAccount("MyService")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	if foundAccount.Name != "MyService" {
		t.Errorf("Expected account name 'MyService' but got '%s'", foundAccount.Name)
	}
	if foundAccount.Secret != "SECRET123" {
		t.Errorf("Expected secret 'SECRET123' but got '%s'", foundAccount.Secret)
	}
	if time.Since(foundAccount.CreatedAt) > 2*time.Second {
		t.Errorf("Expected CreatedAt to be recent but got: '%v'", foundAccount.CreatedAt)
	}

	_ = os.Remove(testFile)
}

func TestAddAccount_WithInvalidName(t *testing.T) {
	err := AddAccount("", "SECRET123", "Password123")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}
}

func TestAddAccount_WithInvalidSecret(t *testing.T) {
	err := AddAccount("MyService", "", "Password123")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}
}

func TestAddAccount_WithInvalidPassword(t *testing.T) {
	err := AddAccount("MyService", "SECRET123", "")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}
}

func TestAddAccount_WithDuplicateAccount(t *testing.T) {
	testFile := "test_duplicate.enc"
	storage.SetStorageFile(testFile)
	_ = os.Remove(testFile)

	err := AddAccount("DupService", "SECRET123", "Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	err = AddAccount("DupService", "SECRET123", "Password123")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	_ = os.Remove(testFile)
}
