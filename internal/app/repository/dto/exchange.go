package dto

type Exchange struct {
	User1ID string `db:"user_id_1"`
	User2ID string `db:"user_id_2"`
	EventID string `db:"event_id"`
}
