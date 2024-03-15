package event

import (
	"context"

	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
)

type IRepository interface {
	Create(ctx context.Context, event *entity.Event) error
	Join(ctx context.Context, eventID string, userID string) error
}
