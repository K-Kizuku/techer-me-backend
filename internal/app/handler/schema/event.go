package schema

type CreateEventInput struct {
	Name       string `json:"name"`
	StartedAt  string `json:"started_at"`
	FinishedAt string `json:"finished_at"`
	Message    string `json:"message"`
	ImageURL   string `json:"image_url"`
	OwnerID    string `json:"owner_id"`
}

type GetEventDetailByIDOutput struct {
	EventID    string `json:"event_id"`
	Name       string `json:"name"`
	StartedAt  string `json:"started_at"`
	FinishedAt string `json:"finished_at"`
	Message    string `json:"message"`
	ImageURL   string `json:"image_url"`
	OwnerID    string `json:"owner_id"`
}
