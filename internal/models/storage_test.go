package models

import (
	"errors"
	"testing"
	"time"
)

func TestInit_WithUninitializedAccounts(t *testing.T) {
	payload := StorageData{}

	payload.Init()

	if payload.Accounts == nil {
		t.Errorf("Expected Accounts to be initialized, but got nil")
	}
}

func TestExists_WithExistingAccount(t *testing.T) {
	payloadAccountName := "gitlab"
	payload := StorageData{Accounts: map[string]Account{
		payloadAccountName: {Name: payloadAccountName, Secret: "ABCDEFGHIJ", CreatedAt: time.Now()},
	}}

	got := payload.Exists(payloadAccountName)

	if !got {
		t.Errorf("Expected true but got false")
	}
}

func TestExists_WithNonExistingAccount(t *testing.T) {
	payloadAccountName := "gitlab"
	payload := StorageData{}

	got := payload.Exists(payloadAccountName)

	if got {
		t.Errorf("Expected false but got true")
	}
}

func TestAddAccount_WithValidAccount(t *testing.T) {
	payloadAccount := Account{Name: "gitlab", Secret: "ABCDEFGHIJ", CreatedAt: time.Now()}
	payload := StorageData{}

	err := payload.AddAccount(payloadAccount)

	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	if _, exists := payload.Accounts[payloadAccount.Name]; !exists {
		t.Fatalf("Expected account '%s' to be added, but it was not found", payloadAccount.Name)
	}
}

func TestAddAccount_WithDuplicateAccount(t *testing.T) {
	payloadAccount := Account{Name: "gitlab", Secret: "ABCDEFGHIJ", CreatedAt: time.Now()}
	payload := StorageData{Accounts: map[string]Account{
		payloadAccount.Name: payloadAccount,
	}}
	expected := errors.New("uma conta com este nome já existe")

	got := payload.AddAccount(payloadAccount)

	if got == nil || got.Error() != expected.Error() {
		t.Errorf("Expected: '%v' but got: '%v'", expected, got)
	}
}

func TestGetAccount_WithExistingAccount(t *testing.T) {
	payloadAccount := Account{Name: "gitlab", Secret: "ABCDEFGHIJ", CreatedAt: time.Now()}
	payload := StorageData{Accounts: map[string]Account{
		payloadAccount.Name: payloadAccount,
	}}

	got, err := payload.GetAccount(payloadAccount.Name)

	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	if got.Name != payloadAccount.Name {
		t.Fatalf("Expected account name '%s' but got '%s'", payloadAccount.Name, got.Name)
	}
}

func TestGetAccount_WithNonExistingAccount(t *testing.T) {
	payload := StorageData{}
	payloadAccountName := "gitlab"
	expected := errors.New("conta não encontrada")

	_, got := payload.GetAccount(payloadAccountName)

	if got == nil || got.Error() != expected.Error() {
		t.Errorf("Expected: '%v' but got: '%v'", expected, got)
	}
}

func TestDeleteAccount_WithExistingAccount(t *testing.T) {
	payloadAccount := Account{Name: "gitlab", Secret: "ABCDEFGHIJ", CreatedAt: time.Now()}
	payload := StorageData{Accounts: map[string]Account{
		payloadAccount.Name: payloadAccount,
	}}

	err := payload.DeleteAccount(payloadAccount.Name)

	if err != nil {
		t.Fatalf("Expected nil but got: '%v'", err)
	}

	if _, exists := payload.Accounts[payloadAccount.Name]; exists {
		t.Fatalf("Expected account '%s' to be deleted, but it still exists", payloadAccount.Name)
	}
}

func TestDeleteAccount_WithNonExistingAccount(t *testing.T) {
	payload := StorageData{}
	payloadAccountName := "gitlab"
	expected := errors.New("conta não encontrada")

	got := payload.DeleteAccount(payloadAccountName)

	if got == nil || got.Error() != expected.Error() {
		t.Errorf("Expected: '%v' but got: '%v'", expected, got)
	}
}
