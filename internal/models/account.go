package models

import "time"

type Account struct {
	Name      string    `json:"name"`
	Secret    string    `json:"secret"`
	CreatedAt time.Time `json:"created_at"`
}
