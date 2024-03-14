package schema

type CreateExchangeInput struct {
	User1ID string `json:"user_id_1"`
	User2ID string `json:"user_id_2"`
	EventID string `json:"event_id"`
}
