package utils

import (
	"errors"
	"testing"
)

func TestValidateServiceName_WithValidName(t *testing.T) {
	payload := "gitlab"

	if got := ValidateServiceName(payload); got != nil {
		t.Errorf("Expected nil but got: '%v'", got)
	}
}

func TestValidateServiceName_WithNameTooShort(t *testing.T) {
	payload := "gi"
	expected := errors.New("o nome do serviço deve ter pelo menos 3 caracteres")

	if got := ValidateServiceName(payload); got.Error() != expected.Error() {
		t.Errorf("Expected: '%v' but got: '%v'", expected, got)
	}
}

func TestValidateServiceName_WithInvalidName(t *testing.T) {
	payload := "gitlab!@#$"
	expected := errors.New("o nome do serviço só pode conter letras, números e hífens")

	if got := ValidateServiceName(payload); got.Error() != expected.Error() {
		t.Errorf("Expected: '%v' but got: '%v'", expected, got)
	}
}

func TestValidateServiceSecret_WithValidSecret(t *testing.T) {
	payload := "ABCDEFGH"

	if got := ValidateServiceSecret(payload); got != nil {
		t.Errorf("Expected nil but got: '%v'", got)
	}
}

func TestValidateServiceSecret_WithEmptySecret(t *testing.T) {
	payload := ""
	expected := errors.New("o secret não pode ser vazio")

	if got := ValidateServiceSecret(payload); got.Error() != expected.Error() {
		t.Errorf("Expected: '%v' but got: '%v'", expected, got)
	}
}

func TestValidatePassword_WithValidPassword(t *testing.T) {
	payload := "12345678"

	if got := ValidatePassword(payload); got != nil {
		t.Errorf("Expected nil but got: '%v'", got)
	}
}

func TestValidatePassword_WithPasswordTooShort(t *testing.T) {
	payload := "123456"
	expected := errors.New("a senha deve ter pelo menos 8 caracteres")

	if got := ValidatePassword(payload); got.Error() != expected.Error() {
		t.Errorf("Expected: '%v' but got: '%v'", expected, got)
	}
}
