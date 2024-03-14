package event

import (
	"context"

	"github.com/K-Kizuku/techer-me-backend/internal/domain/repository/event"
)

type IEventService interface {
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

func (s *Service) Join(ctx context.Context, eventID string, userID string) error {
	if err := s.eventRepo.Join(ctx, eventID, userID); err != nil {
		return err
	}
	return nil
}
