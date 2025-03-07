package account

import (
	"os"
	"testing"
	"time"

	"github.com/JVitoroliv3ira/termotp/internal/storage"
)

func TestRenameAccount_WithValidData(t *testing.T) {
	testFile := "test_rename.enc"
	storage.SetStorageFile(testFile)
	_ = os.Remove(testFile)

	err := AddAccount("OldService", "SECRET123", "Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	err = RenameAccount("OldService", "NewService", "Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	accounts, err := storage.LoadEncrypted("Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	_, err = accounts.GetAccount("OldService")
	if err == nil {
		t.Errorf("Expected an error for 'OldService' but got nil")
	}

	foundAccount, err := accounts.GetAccount("NewService")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	if foundAccount.Name != "NewService" {
		t.Errorf("Expected account name 'NewService' but got '%s'", foundAccount.Name)
	}
	if foundAccount.Secret != "SECRET123" {
		t.Errorf("Expected secret 'SECRET123' but got '%s'", foundAccount.Secret)
	}
	if time.Since(foundAccount.CreatedAt) > 2*time.Second {
		t.Errorf("Expected CreatedAt to be recent but got: '%v'", foundAccount.CreatedAt)
	}

	_ = os.Remove(testFile)
}

func TestRenameAccount_WithShortOldName(t *testing.T) {
	testFile := "test_rename_short.enc"
	storage.SetStorageFile(testFile)
	_ = os.Remove(testFile)

	err := RenameAccount("ab", "NewName", "Password123")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	expectedErr := "informe corretamente o nome da conta que deseja renomear (mínimo 3 caracteres)"
	if err.Error() != expectedErr {
		t.Errorf("Expected error '%s' but got '%s'", expectedErr, err.Error())
	}

	_ = os.Remove(testFile)
}

func TestRenameAccount_WithInvalidNewName(t *testing.T) {
	testFile := "test_rename_invalid.enc"
	storage.SetStorageFile(testFile)
	_ = os.Remove(testFile)

	err := AddAccount("OldService", "SECRET123", "Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	err = RenameAccount("OldService", "", "Password123")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	_ = os.Remove(testFile)
}

func TestRenameAccount_WithShortPassword(t *testing.T) {
	testFile := "test_rename_short_pass.enc"
	storage.SetStorageFile(testFile)
	_ = os.Remove(testFile)

	err := RenameAccount("OldService", "NewService", "short")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	expectedErr := "a senha deve ter pelo menos 8 caracteres"
	if err.Error() != expectedErr {
		t.Errorf("Expected error '%s' but got '%s'", expectedErr, err.Error())
	}

	_ = os.Remove(testFile)
}

func TestRenameAccount_WithNonExistentAccount(t *testing.T) {
	testFile := "test_rename_not_found.enc"
	storage.SetStorageFile(testFile)
	_ = os.Remove(testFile)

	err := RenameAccount("NonExistent", "NewName", "Password123")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	_ = os.Remove(testFile)
}

func TestRenameAccount_WithIncorrectPassword(t *testing.T) {
	testFile := "test_rename_wrong_pass.enc"
	storage.SetStorageFile(testFile)
	_ = os.Remove(testFile)

	err := AddAccount("OldService", "SECRET123", "Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	err = RenameAccount("OldService", "NewService", "WrongPassword")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	_ = os.Remove(testFile)
}

func TestRenameAccount_ToExistingAccount(t *testing.T) {
	testFile := "test_rename_existing.enc"
	storage.SetStorageFile(testFile)
	_ = os.Remove(testFile)

	err := AddAccount("ServiceA", "SECRET123", "Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	err = AddAccount("ServiceB", "SECRET456", "Password123")
	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	err = RenameAccount("ServiceA", "ServiceB", "Password123")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	_ = os.Remove(testFile)
}
