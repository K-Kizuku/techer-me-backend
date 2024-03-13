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
	urls, err := json.Marshal(user.URLs)
	if err != nil {
		return err
	}

	skills, err := json.Marshal(user.Skills)
	if err != nil {
		return err
	}

	var message sql.Null[string]
	message.V = user.Message

	u := &dto.User{
		UserID:      user.ID,
		Name:        user.Name,
		IsOrganizer: user.IsOrganizer,
		ImageURL:    user.ImageURL,
		Message:     message,
		Skills:      string(skills),
		URLs:        string(urls),
	}

	_, err = r.conn.NamedExecContext(ctx, `
		INSERT INTO user_detail (user_id, name, is_organizer, image_url, message, skills, urls)
		VALUES (:user_id, :name, :is_organizer, :image_url, :message, :skills, :urls)
	`, u)

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) SelectByID(ctx context.Context, userID string) (*entity.User, error) {
	return nil, nil
}
func (r *repository) Update(ctx context.Context, user *entity.User) error {
	return nil
}
