package event

import (
	"context"
	"database/sql"

	"github.com/K-Kizuku/techer-me-backend/internal/app/repository/dto"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/repository/event"
	"github.com/K-Kizuku/techer-me-backend/pkg/errors"
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

func (r *repository) Create(ctx context.Context, event *entity.Event) error {

	var message sql.Null[string]
	message.V = event.Message
	e := &dto.Event{
		EventID:    event.ID,
		Name:       event.Name,
		StartedAt:  event.StartedAt,
		FinishedAt: event.FinishedAt,
		Message:    message,
		OwnerID:    event.OwnerID,
		ImageURL:   event.ImageURL,
	}

	_, err := r.conn.ExecContext(ctx, `
		INSERT INTO events (event_id)
		VALUES (?)`, e.EventID)
	if err != nil {
		return errors.HandleError(err)
	}

	_, err = r.conn.NamedExecContext(ctx, `
		INSERT INTO event_details (event_id, name, started_at, finished_at, message, owner_id, image_url)
		VALUES (:event_id, :name, :started_at, :finished_at, :message, :owner_id, :image_url)
	`, e)

	if err != nil {
		return errors.HandleError(err)
	}
	return nil
}

func (r *repository) SelectByID(ctx context.Context, eventID string) (*entity.Event, error) {
	var e dto.Event
	if err := r.conn.QueryRowxContext(ctx, `
		SELECT event_id, name, started_at, finished_at, message, owner_id, image_url
		FROM event_details
		WHERE event_id = ?
	`, eventID).StructScan(&e); err != nil {
		return nil, errors.HandleError(err)
	}
	var eventMessage string
	if e.Message.Valid {
		eventMessage = e.Message.V
	} else {
		eventMessage = ""
	}
	event := entity.Event{
		ID:         e.EventID,
		Name:       e.Name,
		StartedAt:  e.StartedAt,
		FinishedAt: e.FinishedAt,
		Message:    eventMessage,
		OwnerID:    e.OwnerID,
		ImageURL:   e.ImageURL,
	}

	return &event, nil
}
