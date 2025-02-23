package utils

import (
	"errors"
	"regexp"
)

func ValidateServiceName(name string) error {
	if len(name) < 3 {
		return errors.New("o nome do serviço deve ter pelo menos 3 caracteres")
	}

	validName := regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	if !validName.MatchString(name) {
		return errors.New("o nome do serviço só pode conter letras, números e hífens")
	}

	return nil
}

func ValidateServiceSecret(secret string) error {
	if len(secret) == 0 {
		return errors.New("o secret não pode ser vazio")
	}

	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("a senha deve ter pelo menos 8 caracteres")
	}

	return nil
}
