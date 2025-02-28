package storage

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/JVitoroliv3ira/termotp/internal/models"
	"golang.org/x/crypto/argon2"
)

func getStoragePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "secrets.enc"
	}

	var configDir string

	if runtime.GOOS == "windows" {
		appData := os.Getenv("APPDATA")
		if appData == "" {
			appData = homeDir
		}
		configDir = filepath.Join(appData, "TermOTP")
	} else {
		configDir = filepath.Join(homeDir, ".config", "termotp")
	}

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "secrets.enc"
	}

	return filepath.Join(configDir, "secrets.enc")
}

var storageFile = getStoragePath()

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

	if _, exists := storageData.Accounts[account.Name]; exists {
		return errors.New("uma conta com este nome já existe")
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

func SaveAccounts(storageData models.StorageData, password string) error {
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

func GetAccount(name, password string) (models.Account, error) {
	storageData, err := LoadAccounts(password)

	if err != nil {
		return models.Account{}, err
	}

	account, exists := storageData.Accounts[name]
	if !exists {
		return models.Account{}, errors.New("conta não encontrada")
	}

	return account, nil
}

func DeleteAccount(name, password string) error {
	storageData, err := LoadAccounts(password)

	if err != nil {
		return err
	}

	_, exists := storageData.Accounts[name]

	if !exists {
		return errors.New("conta não encontrada")
	}

	delete(storageData.Accounts, name)

	return SaveAccounts(storageData, password)
}
