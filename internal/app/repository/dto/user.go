package dto

import "database/sql"

type User struct {
	ID          int              `db:"id"`
	UserID      string           `db:"user_id"`
	Name        string           `db:"name"`
	IsOrganizer bool             `db:"is_organizer"`
	ImageURL    string           `db:"image_url"`
	Message     sql.Null[string] `db:"message"`
	Skills      string           `db:"skills"`
	URLs        string           `db:"urls"`
}
