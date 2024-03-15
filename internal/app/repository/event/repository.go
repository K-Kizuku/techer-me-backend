package event

import (
	"context"
	"database/sql"

	"github.com/K-Kizuku/techer-me-backend/internal/app/repository/dto"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/repository/event"
	"github.com/K-Kizuku/techer-me-backend/pkg/errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func New(conn *sqlx.DB) event.IRepository {
	return &repository{
		conn: conn,
	}
}

func (r *repository) Create(ctx context.Context, event *entity.Event) error {
	uuid := uuid.New()
	event.ID = uuid.String()

	var message sql.Null[string]
	message.V = event.Message

	e := &dto.Event{
		EventID:    event.ID,
		Name:       event.Name,
		OwnerID:    event.OwnerID,
		StartedAt:  event.StartedAt,
		FinishedAt: event.FinishedAt,
		Message:    message,
		ImageURL:   event.ImageURL,
	}

	_, err := r.conn.NamedExecContext(ctx, `
		INSERT INTO event (event_id)
		VALUES (:event_id)
	`, e)
	if err != nil {
		return errors.HandleError(err)
	}

	_, err = r.conn.NamedExecContext(ctx, `
		INSERT INTO event_details (event_id, name, owner_id, started_at, finished_at, message, image_url)
		VALUES (:event_id, :name, :owner_id, :started_at, :finished_at, :message, :image_url)
	`, e)
	if err != nil {
		return errors.HandleError(err)
	}

	return nil
}

func (r *repository) Join(ctx context.Context, eventID string, userID string) error {
	_, err := r.conn.ExecContext(ctx, `
		INSERT INTO participants (event_id, user_id)
		VALUES (?, ?)
	`, eventID, userID)
	if err != nil {
		return errors.HandleError(err)
	}
	return nil
}
