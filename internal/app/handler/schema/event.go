package schema

type CreateEventInput struct {
	EventID    string `json:"event_id"`
	OwnerID    string `json:"owner_id"`
	Name       string `json:"name"`
	StartedAt  string `json:"started_at"`
	FinishedAt string `json:"finished_at"`
	Message    string `json:"message"`
	ImageURL   string `json:"image_url"`
}

type Event struct {
	EventID    string `json:"event_id"`
	OwnerID    string `json:"owner_id"`
	Name       string `json:"name"`
	StartedAt  string `json:"started_at"`
	FinishedAt string `json:"finished_at"`
	Message    string `json:"message"`
	ImageURL   string `json:"image_url"`
}
