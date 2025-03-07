package models

import "errors"

type StorageData struct {
	Accounts map[string]Account `json:"accounts"`
}

func (s *StorageData) Init() {
	if s.Accounts == nil {
		s.Accounts = make(map[string]Account)
	}
}

func (s *StorageData) Exists(name string) bool {
	_, exists := s.Accounts[name]

	return exists
}

func (s *StorageData) AddAccount(account Account) error {
	if s.Accounts == nil {
		s.Accounts = make(map[string]Account)
	}

	if s.Exists(account.Name) {
		return errors.New("uma conta com este nome já existe")
	}

	s.Accounts[account.Name] = account

	return nil
}

func (s *StorageData) GetAccount(name string) (*Account, error) {
	if !s.Exists(name) {
		return &Account{}, errors.New("conta não encontrada")
	}

	account := s.Accounts[name]
	return &account, nil
}

func (s *StorageData) DeleteAccount(name string) error {
	if !s.Exists(name) {
		return errors.New("conta não encontrada")
	}

	delete(s.Accounts, name)
	return nil
}
