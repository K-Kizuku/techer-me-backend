package user

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/K-Kizuku/techer-me-backend/internal/app/repository/dto"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/repository/user"
	"github.com/K-Kizuku/techer-me-backend/pkg/errors"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func New(conn *sqlx.DB) user.IRepository {
	return &repository{
		conn: conn,
	}
}

func (r *repository) Create(ctx context.Context, id string) error {
	_, err := r.conn.ExecContext(ctx, `
        INSERT INTO users (user_id)
        VALUES (?)
    `, id)

	if err != nil {
		return errors.HandleError(err)
	}
	return nil
}

func (r *repository) CreateDetail(ctx context.Context, user *entity.User) error {
	urls := make(map[string]string)
	urlsJSON, err := json.Marshal(urls)
	if err != nil {
		return errors.New(http.StatusInternalServerError, err)
	}

	skills := make(map[string]string)
	skillsJSON, err := json.Marshal(skills)
	if err != nil {
		return errors.New(http.StatusInternalServerError, err)
	}

	var message sql.Null[string]
	message.V = user.Message

	u := &dto.User{
		UserID:   user.ID,
		Name:     user.Name,
		ImageURL: user.ImageURL,
		Message:  message,
		Skills:   string(skillsJSON),
		URLs:     string(urlsJSON),
	}

	_, err = r.conn.NamedExecContext(ctx, `
		INSERT INTO user_details (user_id, name, image_url, message, skills, urls)
		VALUES (:user_id, :name, :image_url, :message, :skills, :urls)
	`, u)

	if err != nil {
		return errors.HandleError(err)
	}
	return nil
}

func (r *repository) SelectByID(ctx context.Context, userID string) (*entity.User, error) {
	var u dto.User
	if err := r.conn.QueryRowxContext(ctx, `
			SELECT user_id, name, is_organizer, image_url, urls, skills, message FROM user_details WHERE user_id = ?
		`, userID).StructScan(&u); err != nil {
		return nil, errors.HandleError(err)
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
		userMessage = ""
	} else {
		userMessage = u.Message.V
	}

	events := make([]entity.Event, 0)

	rows, err := r.conn.QueryxContext(ctx, `
	SELECT event_id, owner_id, started_at, finished_at, message, image_url FROM event_details JOIN participants USING(event_id) WHERE user_id = ?;
	`, userID)
	if err != nil {
		return nil, errors.HandleError(err)
	}
	for rows.Next() {
		var e dto.Event
		if err := rows.StructScan(&e); err != nil {
			return nil, errors.New(http.StatusInternalServerError, err)
		}
		var eventMessage string
		if e.Message.Valid {
			eventMessage = ""
		} else {
			eventMessage = e.Message.V
		}
		events = append(events, entity.Event{
			ID:         e.EventID,
			Name:       e.Name,
			OwnerID:    e.OwnerID,
			StartedAt:  e.StartedAt,
			FinishedAt: e.FinishedAt,
			Message:    eventMessage,
			ImageURL:   e.ImageURL,
		})
	}

	user := &entity.User{
		ID:          u.UserID,
		Name:        u.Name,
		IsOrganizer: u.IsOrganizer,
		ImageURL:    u.ImageURL,
		Message:     userMessage,
		Skills:      skills,
		URLs:        urls,
		Events:      events,
	}

	return user, nil
}
func (r *repository) Update(ctx context.Context, user *entity.User) error {
	return nil
}
