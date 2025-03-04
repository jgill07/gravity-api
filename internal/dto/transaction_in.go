package dto

import (
	"github.com/google/uuid"
	"github.com/jgill07/gravity-api/internal/models"
)

type TransactionIn struct {
	Type        models.TransactionType `json:"type"`
	Description string                 `json:"description"`
	Amount      float64                `json:"amount"`
}

func (t *TransactionIn) ToModel(uid uuid.UUID) *models.Transaction {
	return &models.Transaction{
		ID:          uid,
		Type:        t.Type,
		Description: t.Description,
		Amount:      t.Amount,
	}
}
