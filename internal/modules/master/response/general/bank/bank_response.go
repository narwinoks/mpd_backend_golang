package bank

import (
	"time"
)

type BankResponse struct {
	ID        string    `json:"id"`
	Bank      string    `json:"bank"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
