package user

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/K-Kizuku/techer-me-backend/internal/app/repository/dto"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/repository/user"
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
		return err
	}
	return nil
}

func (r *repository) CreateDetail(ctx context.Context, user *entity.User) error {
	urls := "{}"
	skills := "{}"

	var message sql.Null[string]
	message.V = user.Message

	u := &dto.User{
		UserID:   user.ID,
		Name:     user.Name,
		ImageURL: user.ImageURL,
		Message:  message,
		Skills:   string(skills),
		URLs:     string(urls),
	}

	_, err := r.conn.NamedExecContext(ctx, `
        INSERT INTO user_details (user_id, name, image_url, message, skills, urls)
        VALUES (:user_id, :name, :image_url, :message, :skills, :urls)
    `, u)

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) SelectByID(ctx context.Context, userID string) (*entity.User, error) {
	var u dto.User
	if err := r.conn.QueryRowxContext(ctx, `
			SELECT user_id, name, is_organizer, image_url, urls, skills, message FROM user_details WHERE user_id = ?
		`, userID).StructScan(&u); err != nil {
		return nil, err
	}

	urls := make(map[entity.URLs]string)
	if err := json.Unmarshal([]byte(u.URLs), &urls); err != nil {
		return nil, err
	}

	skills := make(map[string]string)
	if err := json.Unmarshal([]byte(u.Skills), &skills); err != nil {
		return nil, err
	}

	var message string
	if u.Message.Valid {
		message = ""
	} else {
		message = u.Message.V
	}

	user := &entity.User{
		ID:          u.UserID,
		Name:        u.Name,
		IsOrganizer: u.IsOrganizer,
		ImageURL:    u.ImageURL,
		Message:     message,
		Skills:      skills,
		URLs:        urls,
	}

	return user, nil
}
func (r *repository) Update(ctx context.Context, user *entity.User) error {
	return nil
}
