package storage

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/JVitoroliv3ira/termotp/internal/models"
	"golang.org/x/crypto/argon2"
)

const storageFile = "secrets.enc"

func deriveKey(password string) []byte {
	salt := []byte("static-salt")
	return argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
}

func Encrypt(data []byte, password string) ([]byte, error) {
	key := deriveKey(password)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func Decrypt(encryptedData []byte, password string) ([]byte, error) {
	key := deriveKey(password)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(encryptedData) < nonceSize {
		return nil, errors.New("dados criptografados inválidos")
	}

	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]
	plainData, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.New("senha incorreta ou arquivo corrompido")
	}

	return plainData, nil
}

func SaveAccount(account models.Account, password string) error {
	var storageData models.StorageData

	if _, err := os.Stat(storageFile); err == nil {
		encryptedData, err := os.ReadFile(storageFile)
		if err != nil {
			return err
		}

		decryptedData, err := Decrypt(encryptedData, password)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(decryptedData, &storageData); err != nil {
			return err
		}
	} else {
		storageData = models.StorageData{Accounts: make(map[string]models.Account)}
	}

	storageData.Accounts[account.Name] = account

	jsonData, err := json.Marshal(storageData)
	if err != nil {
		return err
	}

	encryptedData, err := Encrypt(jsonData, password)
	if err != nil {
		return err
	}

	return os.WriteFile(storageFile, encryptedData, 0644)
}

func LoadAccounts(password string) (models.StorageData, error) {
	var storageData models.StorageData

	if _, err := os.Stat(storageFile); err != nil {
		return storageData, errors.New("nenhuma conta cadastrada")
	}

	encryptedData, err := os.ReadFile(storageFile)
	if err != nil {
		return storageData, err
	}

	decryptedData, err := Decrypt(encryptedData, password)
	if err != nil {
		return storageData, errors.New("senha incorreta ou arquivo corrompido")
	}

	if err := json.Unmarshal(decryptedData, &storageData); err != nil {
		return storageData, err
	}

	return storageData, nil
}
