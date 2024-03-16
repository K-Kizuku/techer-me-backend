package event

import (
	"context"

	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
)

type IRepository interface {
	Join(ctx context.Context, eventID string, userID string) error
	Create(ctx context.Context, event *entity.Event) (string, error)
	SelectByID(ctx context.Context, eventID string) (*entity.Event, error)
}
