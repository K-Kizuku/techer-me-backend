package exchange

import (
	"context"

	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
)

type IRepository interface {
	Create(ctx context.Context, exchange entity.Exchange) error
	SelectAllByUserID(ctx context.Context, userID string) ([]entity.ExchangeUser, error)
}
