package event

import (
	"context"
)

type IRepository interface {
	Join(ctx context.Context, eventID string, userID string) error
}
