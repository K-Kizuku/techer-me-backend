package user

import (
	"context"

	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
)

type IRepository interface {
	Create(ctx context.Context, id string) error
	CreateDetail(ctx context.Context, user *entity.User) error
	SelectByID(ctx context.Context, userID string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	SelectEventByID(ctx context.Context, userID string) ([]entity.Event, error)
}
