package event

import (
	"context"

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
