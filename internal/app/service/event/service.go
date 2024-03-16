package event

import (
	"context"
	"net/http"

	"github.com/K-Kizuku/techer-me-backend/internal/app/handler/schema"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/repository/event"
	"github.com/K-Kizuku/techer-me-backend/pkg/errors"
	"github.com/google/uuid"
)

type IEventService interface {
	Join(ctx context.Context, eventID string, userID string) error
	Create(ctx context.Context, input *schema.CreateEventInput) error
	SelectByID(ctx context.Context, eventID string) (*schema.GetEventDetailByIDOutput, error)
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

func (s *Service) Create(ctx context.Context, input *schema.CreateEventInput) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return errors.New(http.StatusInternalServerError, err)
	}
	eventID := uuid.String()
	event := &entity.Event{
		ID:         eventID,
		Name:       input.Name,
		StartedAt:  input.StartedAt,
		FinishedAt: input.FinishedAt,
		Message:    input.Message,
		OwnerID:    input.OwnerID,
		ImageURL:   input.ImageURL,
	}
	if err := s.eventRepo.Create(ctx, event); err != nil {
		return err
	}
	return nil
}

func (s *Service) SelectByID(ctx context.Context, eventID string) (*schema.GetEventDetailByIDOutput, error) {
	event, err := s.eventRepo.SelectByID(ctx, eventID)
	if err != nil {
		return nil, err
	}
	output := &schema.GetEventDetailByIDOutput{
		EventID:    event.ID,
		Name:       event.Name,
		StartedAt:  event.StartedAt,
		FinishedAt: event.FinishedAt,
		Message:    event.Message,
		OwnerID:    event.OwnerID,
		ImageURL:   event.ImageURL,
	}
	return output, nil
}
