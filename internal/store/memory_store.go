package store

import (
	"context"
	"sync"

	"github.com/jgill07/gravity-api/internal/models"
)

var _ TransactionStore = (*MemoryStore)(nil)

// Memory store holds transactions in memory safely
type MemoryStore struct {
	mu           sync.RWMutex
	transactions []*models.Transaction
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		transactions: []*models.Transaction{},
	}
}

func (ms *MemoryStore) Create(ctx context.Context, t *models.Transaction) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.transactions = append(ms.transactions, t)
	return nil
}

func (ms *MemoryStore) Get(ctx context.Context) ([]*models.Transaction, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	return ms.transactions, nil
}
