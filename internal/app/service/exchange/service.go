package exchange

import (
	"context"

	"github.com/K-Kizuku/techer-me-backend/internal/app/handler/schema"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/repository/exchange"
)

type IExchangeService interface {
	Create(ctx context.Context, input *schema.CreateExchangeInput) error
	GetByID(ctx context.Context, userID string) ([]entity.Exchange, error)
}

type Service struct {
	exchangeRepo exchange.IRepository
}

func New(exchangeRepo exchange.IRepository) IExchangeService {
	return &Service{
		exchangeRepo: exchangeRepo,
	}
}

func (s *Service) Create(ctx context.Context, input *schema.CreateExchangeInput) error {
	exchange := entity.Exchange{
		User1ID: input.User1ID,
		User2ID: input.User2ID,
		EventID: input.EventID,
	}
	if err := s.exchangeRepo.Create(ctx, exchange); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetByID(ctx context.Context, userID string) ([]entity.Exchange, error) {
	exchanges, err := s.exchangeRepo.SelectAllByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return exchanges, nil
}
