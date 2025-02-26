package storage

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/JVitoroliv3ira/termotp/internal/models"
)

func TestDeriveKey_WithSamePassword(t *testing.T) {
	password := "password123"
	key1 := deriveKey(password)
	key2 := deriveKey(password)
	if !reflect.DeepEqual(key1, key2) {
		t.Errorf("Expected same key for the same password, got different:\nkey1 = %v\nkey2 = %v", key1, key2)
	}
}

func TestDeriveKey_WithDifferentPasswords(t *testing.T) {
	key1 := deriveKey("password123")
	key2 := deriveKey("another-password")
	if reflect.DeepEqual(key1, key2) {
		t.Errorf("Expected different keys for different passwords:\nkey1 = %v\nkey2 = %v", key1, key2)
	}
}

func TestEncryptDecrypt_WithCorrectPassword(t *testing.T) {
	password := "my-secret-password"
	plaintext := []byte("this is a secret")

	encrypted, err := Encrypt(plaintext, password)
	if err != nil {
		t.Fatalf("Encrypt() error: %v", err)
	}

	decrypted, err := Decrypt(encrypted, password)
	if err != nil {
		t.Fatalf("Decrypt() error: %v", err)
	}

	if !reflect.DeepEqual(decrypted, plaintext) {
		t.Errorf("Expected %q, got %q", plaintext, decrypted)
	}
}

func TestEncryptDecrypt_WithWrongPassword(t *testing.T) {
	password := "correct-password"
	plaintext := []byte("this is a secret")

	encrypted, err := Encrypt(plaintext, password)
	if err != nil {
		t.Fatalf("Encrypt() error: %v", err)
	}

	_, err = Decrypt(encrypted, "wrong-password")
	if err == nil {
		t.Fatalf("Expected error when decrypting with wrong password, got none")
	}
}

func TestSaveAccountAndLoadAccounts(t *testing.T) {
	tempDir := t.TempDir()
	originalFile := storageFile
	defer func() { storageFile = originalFile }()
	storageFile = filepath.Join(tempDir, "secrets.enc")

	password := "test-password"
	acc := models.Account{
		Name:   "MyTestAccount",
		Secret: "test-secret",
	}

	if err := SaveAccount(acc, password); err != nil {
		t.Fatalf("SaveAccount() error: %v", err)
	}

	data, err := LoadAccounts(password)
	if err != nil {
		t.Fatalf("LoadAccounts() error: %v", err)
	}

	got, exists := data.Accounts[acc.Name]
	if !exists {
		t.Fatalf("Expected to find account %q, but it wasn't loaded", acc.Name)
	}

	if !reflect.DeepEqual(acc, got) {
		t.Errorf("Expected %v, got %v", acc, got)
	}
}

func TestLoadAccounts_NoFile(t *testing.T) {
	tempDir := t.TempDir()
	originalFile := storageFile
	defer func() { storageFile = originalFile }()
	storageFile = filepath.Join(tempDir, "doesnotexist.enc")

	password := "test-password"
	_, err := LoadAccounts(password)
	if err == nil {
		t.Fatalf("Expected error when loading from non-existent file, got none")
	}
}

func TestGetAccount_WithExistingAndMissing(t *testing.T) {
	tempDir := t.TempDir()
	originalFile := storageFile
	defer func() { storageFile = originalFile }()
	storageFile = filepath.Join(tempDir, "secrets.enc")

	password := "test-password"
	acc1 := models.Account{Name: "Account1", Secret: "secret-1"}
	acc2 := models.Account{Name: "Account2", Secret: "secret-2"}

	if err := SaveAccount(acc1, password); err != nil {
		t.Fatalf("SaveAccount() error: %v", err)
	}
	if err := SaveAccount(acc2, password); err != nil {
		t.Fatalf("SaveAccount() error: %v", err)
	}

	got, err := GetAccount(acc1.Name, password)
	if err != nil {
		t.Fatalf("GetAccount() error: %v", err)
	}
	if !reflect.DeepEqual(acc1, got) {
		t.Errorf("Expected %v, got %v", acc1, got)
	}

	_, err = GetAccount("NoSuchAccount", password)
	if err == nil {
		t.Fatalf("Expected error for non-existent account, got none")
	}
}

func TestLoadAccounts_WithWrongPassword(t *testing.T) {
	tempDir := t.TempDir()
	originalFile := storageFile
	defer func() { storageFile = originalFile }()
	storageFile = filepath.Join(tempDir, "secrets.enc")

	password := "correct-password"
	acc := models.Account{Name: "TestAccount", Secret: "test-secret"}

	if err := SaveAccount(acc, password); err != nil {
		t.Fatalf("SaveAccount() error: %v", err)
	}

	_, err := LoadAccounts("wrong-password")
	if err == nil {
		t.Fatalf("Expected error with wrong password, got none")
	}
}

func TestDecrypt_WithInvalidData(t *testing.T) {
	data := []byte("xyz")
	password := "some-password"

	_, err := Decrypt(data, password)
	if err == nil {
		t.Fatalf("Expected error for invalid data, got none")
	}
}

func TestGetStoragePath_ReturnsNonEmpty(t *testing.T) {
	path := getStoragePath()
	if path == "" {
		t.Errorf("Expected a non-empty string, got %q", path)
	}
}

func TestSaveAccount_FilePermissions(t *testing.T) {
	tempDir := t.TempDir()
	originalFile := storageFile
	defer func() { storageFile = originalFile }()
	storageFile = filepath.Join(tempDir, "secrets.enc")

	password := "test-password"
	acc := models.Account{Name: "PermissionTestAccount", Secret: "secret"}

	if err := SaveAccount(acc, password); err != nil {
		t.Fatalf("SaveAccount() error: %v", err)
	}
	info, err := os.Stat(storageFile)
	if err != nil {
		t.Fatalf("os.Stat() error: %v", err)
	}
	if mode := info.Mode().Perm(); mode != 0644 {
		t.Errorf("Expected permission 0644, got %v", mode)
	}
}
