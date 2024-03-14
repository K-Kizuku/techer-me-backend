package dto

import "database/sql"

type Event struct {
	EventID    string           `db:"event_id"`
	Name       string           `db:"name"`
	OwnerID    string           `db:"owner_id"`
	StartedAt  string           `db:"started_at"`
	FinishedAt string           `db:"finished_at"`
	Message    sql.Null[string] `db:"message"`
	ImageURL   string           `db:"image_url"`
}
