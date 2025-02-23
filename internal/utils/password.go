package utils

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

func PromptPassword() (string, error) {
	fmt.Println("Digite sua senha: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()

	if err != nil {
		return "", err
	}

	return string(bytePassword), nil
}
