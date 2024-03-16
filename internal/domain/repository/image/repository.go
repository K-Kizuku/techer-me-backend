package image

import (
	"context"

	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
)

type IRepository interface {
	GenerateSignedURL(ctx context.Context, image entity.Image) (string, error)
}
