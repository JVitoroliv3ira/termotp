package security

import (
	"bytes"
	"errors"
	"testing"
)

func TestEncryptDecrypt_WithValidData(t *testing.T) {
	payload := []byte("dados secretos")
	password := "senhaSegura"

	encrypted, err := Encrypt(payload, password)
	if err != nil {
		t.Fatalf("Expected nil but got error: '%v'", err)
	}

	decrypted, err := Decrypt(encrypted, password)
	if err != nil {
		t.Fatalf("Expected nil but got error: '%v'", err)
	}

	if !bytes.Equal(payload, decrypted) {
		t.Fatalf("Expected decrypted data to be '%s' but got '%s'", payload, decrypted)
	}
}

func TestDecrypt_WithIncorrectPassword(t *testing.T) {
	payload := []byte("dados secretos")
	password := "senhaCorreta"
	wrongPassword := "senhaErrada"

	encrypted, err := Encrypt(payload, password)
	if err != nil {
		t.Fatalf("Expected nil but got error: '%v'", err)
	}

	_, err = Decrypt(encrypted, wrongPassword)

	expected := errors.New("senha incorreta ou arquivo corrompido")
	if err == nil || err.Error() != expected.Error() {
		t.Errorf("Expected: '%v' but got: '%v'", expected, err)
	}
}

func TestDecrypt_WithCorruptedData(t *testing.T) {
	payload := []byte("dados secretos")
	password := "senhaSegura"

	encrypted, err := Encrypt(payload, password)
	if err != nil {
		t.Fatalf("Expected nil but got error: '%v'", err)
	}

	encrypted[len(encrypted)-1] ^= 0xFF

	_, err = Decrypt(encrypted, password)

	expected := errors.New("senha incorreta ou arquivo corrompido")
	if err == nil || err.Error() != expected.Error() {
		t.Errorf("Expected: '%v' but got: '%v'", expected, err)
	}
}

func TestDecrypt_WithInvalidCiphertext(t *testing.T) {
	password := "senhaSegura"

	_, err := Decrypt([]byte("dados inválidos"), password)

	expected := errors.New("senha incorreta ou arquivo corrompido")
	if err == nil || err.Error() != expected.Error() {
		t.Errorf("Expected: '%v' but got: '%v'", expected, err)
	}
}
