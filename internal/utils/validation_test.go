package utils

import (
	"errors"
	"testing"
)

func TestValidateServiceName_WithValidName(t *testing.T) {
	serviceName := "gitlab"

	if got := ValidateServiceName(serviceName); got != nil {
		t.Errorf("ValidateServiceName() returned error: %v; want nil", got)
	}
}

func TestValidateServiceName_WithNameTooShort(t *testing.T) {
	serviceName := "a"
	want := errors.New("o nome do serviço deve ter pelo menos 3 caracteres")

	if got := ValidateServiceName(serviceName); got.Error() != want.Error() {
		t.Errorf("ValidateServiceName() = %q; want = %q", got, want)
	}
}

func TestValidateServiceName_WithInvalidCharacters(t *testing.T) {
	serviceName := "abc$%¨&*"
	want := errors.New("o nome do serviço só pode conter letras, números e hífens")

	if got := ValidateServiceName(serviceName); got.Error() != want.Error() {
		t.Errorf("ValidateServiceName() = %q; want = %q", got, want)
	}
}

func TestValidateServiceName_WithEmptyName(t *testing.T) {
	serviceName := ""
	want := errors.New("o nome do serviço deve ter pelo menos 3 caracteres")

	if got := ValidateServiceName(serviceName); got.Error() != want.Error() {
		t.Errorf("ValidateServiceName() = %q; want = %q", got, want)
	}
}

func TestValidateServiceSecret_WithValidSecret(t *testing.T) {
	serviceSecret := "ABCDEFGHI"

	if got := ValidateServiceSecret(serviceSecret); got != nil {
		t.Errorf("ValidateServiceSecret() returned error: %v; want nil", got)
	}
}

func TestValidateServiceSecret_WithEmptyName(t *testing.T) {
	serviceSecret := ""
	want := errors.New("o secret não pode ser vazio")

	if got := ValidateServiceSecret(serviceSecret); got.Error() != want.Error() {
		t.Errorf("ValidateServiceSecret() = %q; want = %q", got, want)
	}
}

func TestValidatePassword_WithValidPassword(t *testing.T) {
	password := "12345678"

	if got := ValidatePassword(password); got != nil {
		t.Errorf("ValidatePassword() returned error: %v; want nil", got)
	}
}

func TestValidatePassword_WithPasswordTooShort(t *testing.T) {
	password := "1234567"
	want := errors.New("a senha deve ter pelo menos 8 caracteres")

	if got := ValidatePassword(password); got.Error() != want.Error() {
		t.Errorf("ValidatePassword() = %q; want = %q", got, want)
	}
}
