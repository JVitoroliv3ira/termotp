package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/JVitoroliv3ira/termotp/internal/models"
	"github.com/JVitoroliv3ira/termotp/internal/security"
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

func SaveEncrypted(storageData *models.StorageData, password string) error {
	jsonData, err := json.Marshal(storageData)
	if err != nil {
		return err
	}

	encryptedData, err := security.Encrypt(jsonData, password)
	if err != nil {
		return err
	}

	return os.WriteFile(storageFile, encryptedData, 0644)
}

func LoadEncrypted(password string) (*models.StorageData, error) {
	storageData := &models.StorageData{}
	storageData.Init()

	if _, err := os.Stat(storageFile); errors.Is(err, os.ErrNotExist) {
		return storageData, nil
	}

	encryptedData, err := os.ReadFile(storageFile)
	if err != nil {
		return storageData, err
	}

	decryptedData, err := security.Decrypt(encryptedData, password)
	if err != nil {
		return storageData, err
	}

	if err := json.Unmarshal(decryptedData, &storageData); err != nil {
		return storageData, err
	}

	return storageData, nil
}
