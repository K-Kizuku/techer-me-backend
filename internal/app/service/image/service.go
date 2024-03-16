package image

import (
	"context"

	"github.com/K-Kizuku/techer-me-backend/internal/app/handler/schema"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/repository/image"
)

type IImageService interface {
	GenerateSignedURL(ctx context.Context, input *schema.GenerateImageInput) (*schema.GenerateImageOutput, error)
}

type Service struct {
	imageRepo image.IRepository
}

func New(imageRepo image.IRepository) IImageService {
	return &Service{
		imageRepo: imageRepo,
	}
}

func (s *Service) GenerateSignedURL(ctx context.Context, input *schema.GenerateImageInput) (*schema.GenerateImageOutput, error) {
	image := entity.Image{
		ObjectName: input.ObjectName,
		Type:       input.Type,
	}
	url, err := s.imageRepo.GenerateSignedURL(ctx, image)
	if err != nil {
		return nil, err
	}
	return &schema.GenerateImageOutput{
		URL: url,
	}, nil
}
