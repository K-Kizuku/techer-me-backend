package event

import (
	"context"

	"github.com/K-Kizuku/techer-me-backend/internal/app/handler/schema"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/repository/event"
)

type IEventService interface {
	Create(ctx context.Context, input *schema.CreateEventInput) error
	Join(ctx context.Context, eventID string, userID string) error
}

type Service struct {
	eventRepo event.IRepository
}

func New(eventRepo event.IRepository) IEventService {
	return &Service{
		eventRepo: eventRepo,
	}
}

func (s *Service) Create(ctx context.Context, input *schema.CreateEventInput) error {
	exchange := &entity.Event{
		ID:         input.EventID,
		OwnerID:    input.OwnerID,
		Name:       input.Name,
		StartedAt:  input.StartedAt,
		FinishedAt: input.FinishedAt,
		Message:    input.Message,
		ImageURL:   input.ImageURL,
	}
	if err := s.eventRepo.Create(ctx, exchange); err != nil {
		return err
	}
	return nil
}

func (s *Service) Join(ctx context.Context, eventID string, userID string) error {
	if err := s.eventRepo.Join(ctx, eventID, userID); err != nil {
		return err
	}
	return nil
}
