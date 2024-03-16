package schema

import "github.com/K-Kizuku/techer-me-backend/internal/domain/entity"

type CreateExchangeInput struct {
	User1ID string `json:"user_id_1"`
	User2ID string `json:"user_id_2"`
	EventID string `json:"event_id"`
}

type Exchange struct {
	UserID   string                 `json:"user_id"`
	Name     string                 `json:"name"`
	ImageURL string                 `json:"image_url"`
	Message  string                 `json:"message"`
	Skills   map[string]string      `json:"skills"`
	URLs     map[entity.URLs]string `json:"urls"`
	Times    int                    `json:"times"`
}

type Sticker struct {
	UserID   string  `json:"user_id"`
	ImageURL string  `json:"image_url"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
}

type GetExchangesOutput struct {
	Exchanges []Exchange `json:"exchanges"`
	Stickers  []Sticker  `json:"stickers"`
}
