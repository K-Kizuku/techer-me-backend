package schema

type CreateUserInput struct {
	UserID string `json:"user_id"`
}

type CreateUserDetailInput struct {
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	IsOrganizer bool   `json:"is_organizer"`
	ImageURL    string `json:"image_url"`
	Message     string `json:"message"`
	Skills      string `json:"skills"`
	URLs        string `json:"urls"`
}
