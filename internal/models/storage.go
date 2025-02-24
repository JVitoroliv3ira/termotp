package models

type StorageData struct {
	Accounts map[string]Account `json:"accounts"`
}
