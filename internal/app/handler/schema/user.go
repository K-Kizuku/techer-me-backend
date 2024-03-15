package schema

import "github.com/K-Kizuku/techer-me-backend/internal/domain/entity"

type CreateUserInput struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
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

type GetMeOutput struct {
	UserID      string                 `json:"user_id"`
	Name        string                 `json:"name"`
	IsOrganizer bool                   `json:"is_organizer"`
	ImageURL    string                 `json:"image_url"`
	Message     string                 `json:"message"`
	Skills      map[string]string      `json:"skills"`
	URLs        map[entity.URLs]string `json:"urls"`
	Events      []Event                `json:"events"`
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

type GetByIDInput struct {
	UserID string `json:"user_id"`
}

type GetByIDOutput struct {
	UserID      string                 `json:"user_id"`
	Name        string                 `json:"name"`
	IsOrganizer bool                   `json:"is_organizer"`
	ImageURL    string                 `json:"image_url"`
	Message     string                 `json:"message"`
	Skills      map[string]string      `json:"skills"`
	URLs        map[entity.URLs]string `json:"urls"`
	Events      []Event                `json:"events"`
}

type GetEventByIDOutput struct {
	Events []Event `json:"events"`
}
