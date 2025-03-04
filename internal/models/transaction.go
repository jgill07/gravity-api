package models

import "github.com/google/uuid"

type TransactionType string

const (
	TransactionIncome  TransactionType = "income"
	TransactionExpense TransactionType = "expense"
)

type Transaction struct {
	ID          uuid.UUID       `json:"id"`
	Type        TransactionType `json:"type"`
	Description string          `json:"description"`
	Amount      float64         `json:"amount"`
}
