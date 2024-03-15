package exchange

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/K-Kizuku/techer-me-backend/internal/app/repository/dto"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/repository/exchange"
	"github.com/K-Kizuku/techer-me-backend/pkg/errors"
	"github.com/jmoiron/sqlx"
	"golang.org/x/exp/maps"
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

func (r *repository) SelectAllByUserID(ctx context.Context, userID string) ([]entity.ExchangeUser, error) {
	exchanges := make([]entity.ExchangeUser, 0)

	rows, err := r.conn.QueryxContext(ctx, `
		SELECT user_id_1, user_id_2, event_id FROM exchanges
		WHERE user_id_1 = ? OR user_id_2 = ?
	`, userID, userID)
	if err != nil {
		return nil, errors.HandleError(err)
	}
	userMap := make(map[string]int)
	for rows.Next() {
		var e dto.Exchange
		if err := rows.StructScan(&e); err != nil {
			return nil, errors.New(http.StatusInternalServerError, err)
		}
		if e.User1ID == userID {
			if _, ok := userMap[e.User2ID]; !ok {

				userMap[e.User2ID] = 1
			} else {
				userMap[e.User2ID]++
			}
		} else {
			if _, ok := userMap[e.User1ID]; !ok {
				userMap[e.User1ID] = 1
			} else {
				userMap[e.User1ID]++
			}
		}
	}
	userList := maps.Keys(userMap)
	q, params, err := sqlx.In(`
		SELECT user_id, name, image_url, message, skills, urls FROM user_details WHERE user_id IN (?)
	`, userList)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err)
	}
	rows2, err := r.conn.QueryxContext(ctx, q, params...)
	if err != nil {
		return nil, errors.HandleError(err)
	}
	for rows2.Next() {
		var u dto.User
		if err := rows2.StructScan(&u); err != nil {
			return nil, errors.New(http.StatusInternalServerError, err)
		}
		urls := make(map[entity.URLs]string)
		if err := json.Unmarshal([]byte(u.URLs), &urls); err != nil {
			return nil, errors.New(http.StatusInternalServerError, err)
		}
		skills := make(map[string]string)
		if err := json.Unmarshal([]byte(u.Skills), &skills); err != nil {
			return nil, errors.New(http.StatusInternalServerError, err)
		}
		var userMessage string
		if u.Message.Valid {
			userMessage = u.Message.V
		} else {
			userMessage = ""
		}
		exchanges = append(exchanges, entity.ExchangeUser{
			UserID:   u.UserID,
			Name:     u.Name,
			ImageURL: u.ImageURL,
			Message:  userMessage,
			Skills:   skills,
			URLs:     urls,
			Times:    userMap[u.UserID],
		})
	}
	return exchanges, nil
}
