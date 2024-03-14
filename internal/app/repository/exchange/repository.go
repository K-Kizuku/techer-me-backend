package exchange

import (
	"context"
	"net/http"

	"github.com/K-Kizuku/techer-me-backend/internal/app/repository/dto"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/repository/exchange"
	"github.com/K-Kizuku/techer-me-backend/pkg/errors"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func New(conn *sqlx.DB) exchange.IRepository {
	return &repository{
		conn: conn,
	}
}

func (r *repository) Create(ctx context.Context, exchange entity.Exchange) error {
	user1ID := exchange.User1ID
	user2ID := exchange.User2ID
	if user1ID > user2ID {
		user1ID, user2ID = user2ID, user1ID
	}

	e := &dto.Exchange{
		User1ID: user1ID,
		User2ID: user2ID,
		EventID: exchange.EventID,
	}

	_, err := r.conn.NamedExecContext(ctx, `
		INSERT INTO exchanges (user_id_1, user_id_2, event_id)
		VALUES (:user_id_1, :user_id_2, :event_id)
	`, e)
	if err != nil {
		return errors.HandleError(err)
	}
	return nil
}

func (r *repository) SelectAllByUserID(ctx context.Context, userID string) ([]entity.Exchange, error) {
	exchanges := make([]entity.Exchange, 0)

	rows, err := r.conn.QueryxContext(ctx, `
		SELECT user_id_1, user_id_2, event_id FROM exchanges
		WHERE user_id_1 = $1 OR user_id_2 = $1
	`, userID, userID)
	if err != nil {
		return nil, errors.HandleError(err)
	}
	for rows.Next() {
		var e dto.Exchange
		if err := rows.StructScan(&e); err != nil {
			return nil, errors.New(http.StatusInternalServerError, err)
		}
		exchanges = append(exchanges, entity.Exchange{
			User1ID: e.User1ID,
			User2ID: e.User2ID,
			EventID: e.EventID,
		})
	}
	return exchanges, nil
}
