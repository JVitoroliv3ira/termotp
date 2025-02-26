package storage

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/JVitoroliv3ira/termotp/internal/models"
)

func TestDeriveKey_WithSamePassword(t *testing.T) {
	p := "password123"
	k1 := deriveKey(p)
	k2 := deriveKey(p)
	if !reflect.DeepEqual(k1, k2) {
		t.Errorf("deriveKey() with the same password returned different keys: key1 = %v, key2 = %v", k1, k2)
	}
}

func TestDeriveKey_WithDifferentPasswords(t *testing.T) {
	k1 := deriveKey("password123")
	k2 := deriveKey("another-password")
	if reflect.DeepEqual(k1, k2) {
		t.Errorf("deriveKey() with different passwords returned the same key: %v", k1)
	}
}

func TestEncryptDecrypt_Success(t *testing.T) {
	p := "my-secret"
	d := []byte("top-secret-data")
	enc, err := Encrypt(d, p)
	if err != nil {
		t.Fatalf("Encrypt() returned an unexpected error: %v", err)
	}
	dec, err := Decrypt(enc, p)
	if err != nil {
		t.Fatalf("Decrypt() returned an unexpected error: %v", err)
	}
	if !reflect.DeepEqual(dec, d) {
		t.Errorf("Decrypt() = %q; want %q", dec, d)
	}
}

func TestEncryptDecrypt_WithWrongPassword(t *testing.T) {
	p := "correct-pwd"
	d := []byte("some-data")
	enc, err := Encrypt(d, p)
	if err != nil {
		t.Fatalf("Encrypt() returned an unexpected error: %v", err)
	}
	_, err = Decrypt(enc, "wrong-pwd")
	if err == nil {
		t.Fatalf("Decrypt() should have returned an error for wrong password, but got none")
	}
}

func TestEncryptDecrypt_WithInvalidData(t *testing.T) {
	_, err := Decrypt([]byte("invalid"), "pwd")
	if err == nil {
		t.Fatalf("Decrypt() should have returned an error for invalid data, but got none")
	}
}

func TestSaveAccount_Success(t *testing.T) {
	dir := t.TempDir()
	originalFile := storageFile
	defer func() { storageFile = originalFile }()
	storageFile = filepath.Join(dir, "secrets.enc")

	p := "test-pwd"
	a := models.Account{Name: "TestAccount", Secret: "test-secret"}

	if err := SaveAccount(a, p); err != nil {
		t.Fatalf("SaveAccount() returned an unexpected error: %v", err)
	}
	if _, err := os.Stat(storageFile); err != nil {
		t.Fatalf("os.Stat() should have succeeded, but got error: %v", err)
	}
}

func TestSaveAccount_FilePermissions(t *testing.T) {
	dir := t.TempDir()
	originalFile := storageFile
	defer func() { storageFile = originalFile }()
	storageFile = filepath.Join(dir, "secrets.enc")

	p := "test-pwd"
	a := models.Account{Name: "PermTest", Secret: "perm-secret"}

	if err := SaveAccount(a, p); err != nil {
		t.Fatalf("SaveAccount() returned an unexpected error: %v", err)
	}
	info, err := os.Stat(storageFile)
	if err != nil {
		t.Fatalf("os.Stat() returned an unexpected error: %v", err)
	}
	if mode := info.Mode().Perm(); mode != 0644 {
		t.Errorf("os.Stat() = %v; want 0644", mode)
	}
}

func TestLoadAccounts_Success(t *testing.T) {
	dir := t.TempDir()
	originalFile := storageFile
	defer func() { storageFile = originalFile }()
	storageFile = filepath.Join(dir, "secrets.enc")

	p := "test-pwd"
	a := models.Account{Name: "TestAccount", Secret: "test-secret"}

	if err := SaveAccount(a, p); err != nil {
		t.Fatalf("SaveAccount() returned an unexpected error: %v", err)
	}
	loaded, err := LoadAccounts(p)
	if err != nil {
		t.Fatalf("LoadAccounts() returned an unexpected error: %v", err)
	}
	got, ok := loaded.Accounts[a.Name]
	if !ok {
		t.Fatalf("LoadAccounts() did not return the account: %q", a.Name)
	}
	if !reflect.DeepEqual(a, got) {
		t.Errorf("LoadAccounts() = %v; want %v", got, a)
	}
}

func TestLoadAccounts_WithNoFile(t *testing.T) {
	dir := t.TempDir()
	originalFile := storageFile
	defer func() { storageFile = originalFile }()
	storageFile = filepath.Join(dir, "nofile.enc")

	if _, err := LoadAccounts("some-pwd"); err == nil {
		t.Fatalf("LoadAccounts() should have returned an error for non-existent file, but got none")
	}
}

func TestLoadAccounts_WithWrongPassword(t *testing.T) {
	dir := t.TempDir()
	originalFile := storageFile
	defer func() { storageFile = originalFile }()
	storageFile = filepath.Join(dir, "secrets.enc")

	p := "correct-pwd"
	a := models.Account{Name: "MyAcc", Secret: "abc123"}

	if err := SaveAccount(a, p); err != nil {
		t.Fatalf("SaveAccount() returned an unexpected error: %v", err)
	}
	if _, err := LoadAccounts("wrong-pwd"); err == nil {
		t.Fatalf("LoadAccounts() should have returned an error for wrong password, but got none")
	}
}

func TestGetAccount_Existing(t *testing.T) {
	dir := t.TempDir()
	originalFile := storageFile
	defer func() { storageFile = originalFile }()
	storageFile = filepath.Join(dir, "secrets.enc")

	p := "abc123"
	a := models.Account{Name: "Acc1", Secret: "secret-1"}

	if err := SaveAccount(a, p); err != nil {
		t.Fatalf("SaveAccount() returned an unexpected error: %v", err)
	}
	got, err := GetAccount(a.Name, p)
	if err != nil {
		t.Fatalf("GetAccount() returned an unexpected error: %v", err)
	}
	if !reflect.DeepEqual(a, got) {
		t.Errorf("GetAccount() = %v; want %v", got, a)
	}
}

func TestGetAccount_Missing(t *testing.T) {
	dir := t.TempDir()
	originalFile := storageFile
	defer func() { storageFile = originalFile }()
	storageFile = filepath.Join(dir, "secrets.enc")

	p := "xyz789"
	a := models.Account{Name: "Acc1", Secret: "secret-1"}

	if err := SaveAccount(a, p); err != nil {
		t.Fatalf("SaveAccount() returned an unexpected error: %v", err)
	}
	if _, err := GetAccount("NonExistent", p); err == nil {
		t.Fatalf("GetAccount() should have returned an error for missing account, but got none")
	}
}

func TestGetStoragePath_ReturnsNonEmpty(t *testing.T) {
	if p := getStoragePath(); p == "" {
		t.Errorf("getStoragePath() = %q; want a non-empty path", p)
	}
}
