package utils

import (
	"fmt"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func PromptHiddenInput(prompt string) (string, error) {
	fmt.Print(prompt)
	byteInput, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()

	if err != nil {
		return "", err
	}

	return string(byteInput), nil
}

func PromptPassword() (string, error) {
	return PromptHiddenInput("Digite sua senha: ")
}

func PromptSecret() (string, error) {
	secret, err := PromptHiddenInput("Digite o Secret TOTP: ")
	if err != nil {
		return "", err
	}

	cleanedSecret := strings.ToUpper(strings.ReplaceAll(secret, " ", ""))
	return cleanedSecret, nil
}
