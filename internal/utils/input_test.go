package utils

import (
	"errors"
	"testing"
)

func mockReadPasswordWithInput(input string) func() {
	original := readPasswordFunc
	readPasswordFunc = func(fd int) ([]byte, error) {
		return []byte(input), nil
	}
	return func() {
		readPasswordFunc = original
	}
}

func mockReadPasswordWithError() func() {
	original := readPasswordFunc
	readPasswordFunc = func(fd int) ([]byte, error) {
		return nil, errors.New("simulated read error")
	}
	return func() {
		readPasswordFunc = original
	}
}

func TestPromptHiddenInput_WithValidPassword(t *testing.T) {
	payload := "mySecurePassword"
	expected := "mySecurePassword"

	restore := mockReadPasswordWithInput(payload)
	defer restore()

	got, err := PromptHiddenInput("Digite sua senha: ")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if got != expected {
		t.Errorf("Expected: '%s' but got: '%s'", expected, got)
	}
}

func TestPromptHiddenInput_WithReadError(t *testing.T) {
	restore := mockReadPasswordWithError()
	defer restore()

	_, err := PromptHiddenInput("Digite sua senha: ")

	if err == nil {
		t.Errorf("Expected an error but got nil")
	}
}

func TestPromptPassword_WithValidPassword(t *testing.T) {
	payload := "mySecurePassword"
	expected := "mySecurePassword"

	restore := mockReadPasswordWithInput(payload)
	defer restore()

	got, err := PromptPassword()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if got != expected {
		t.Errorf("Expected: '%s' but got: '%s'", expected, got)
	}
}

func TestPromptSecret_WithValidSecret(t *testing.T) {
	payload := "ABCDEFGHIJK"
	expected := "ABCDEFGHIJK"

	restore := mockReadPasswordWithInput(payload)
	defer restore()

	got, err := PromptSecret()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if got != expected {
		t.Errorf("Expected: '%s' but got: '%s'", expected, got)
	}
}

func TestPromptSecret_WithWhiteSpaces(t *testing.T) {
	payload := "ABCD EFGH IJKL"
	expected := "ABCDEFGHIJKL"

	restore := mockReadPasswordWithInput(payload)
	defer restore()

	got, err := PromptSecret()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if got != expected {
		t.Errorf("Expected: '%s' but got: '%s'", expected, got)
	}
}

func TestPromptSecret_WithWhiteSpacesAndLowerCase(t *testing.T) {
	payload := "abcd efgh ijkl"
	expected := "ABCDEFGHIJKL"

	restore := mockReadPasswordWithInput(payload)
	defer restore()

	got, err := PromptSecret()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if got != expected {
		t.Errorf("Expected: '%s' but got: '%s'", expected, got)
	}
}

func TestPromptSecret_WithReadError(t *testing.T) {
	restore := mockReadPasswordWithError()
	defer restore()

	_, err := PromptSecret()

	if err == nil {
		t.Errorf("Expected an error but got nil")
	}
}
