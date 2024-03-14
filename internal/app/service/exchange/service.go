package exchange

import (
	"context"

	"github.com/K-Kizuku/techer-me-backend/internal/app/handler/schema"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/repository/exchange"
)

type IExchangeService interface {
	Create(ctx context.Context, input *schema.CreateExchangeInput) error
	GetByID(ctx context.Context, userID string) (*schema.GetExchangesOutput, error)
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

func (s *Service) GetByID(ctx context.Context, userID string) (*schema.GetExchangesOutput, error) {
	exchanges, err := s.exchangeRepo.SelectAllByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	output := make([]schema.Exchange, 0)
	for _, exchange := range exchanges {
		if exchange.User1ID == userID {
			output = append(output, schema.Exchange{
				UserID:  exchange.User2ID,
				EventID: exchange.EventID,
			})
		} else {
			output = append(output, schema.Exchange{
				UserID:  exchange.User1ID,
				EventID: exchange.EventID,
			})
		}
	}

	return &schema.GetExchangesOutput{
		Exchanges: output,
	}, nil
}
