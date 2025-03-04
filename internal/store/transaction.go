package store

import (
	"context"

	"github.com/jgill07/gravity-api/internal/models"
)

type TransactionStore interface {
	Create(ctx context.Context, model *models.Transaction) error
	Get(ctx context.Context) ([]*models.Transaction, error)
}
