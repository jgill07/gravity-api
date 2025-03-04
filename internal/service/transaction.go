package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jgill07/gravity-api/internal/config"
	"github.com/jgill07/gravity-api/internal/dto"
	"github.com/jgill07/gravity-api/internal/models"
	"github.com/jgill07/gravity-api/internal/store"
)

type Service struct {
	Config *config.Config
	tStore store.TransactionStore
}

func NewService(cfg *config.Config, tstore store.TransactionStore) *Service {
	return &Service{
		Config: cfg,
		tStore: tstore,
	}
}

func (s *Service) GetTransactions(ctx context.Context) ([]*models.Transaction, error) {
	return s.tStore.Get(ctx)
}

func (s *Service) CreateTransaction(ctx context.Context, model dto.TransactionIn) error {
	return s.tStore.Create(ctx, model.ToModel(uuid.New()))
}
