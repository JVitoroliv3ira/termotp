package utils

import (
	"errors"

	"github.com/atotto/clipboard"
)

var clipboardWriteFunc = clipboard.WriteAll

func CopyToClipboard(text string) error {
	if clipboard.Unsupported {
		return errors.New("não foi possível acessar a área de transferência")
	}

	return clipboardWriteFunc(text)
}
