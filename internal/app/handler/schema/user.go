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

type UpdateUserInput struct {
	Name        string                 `json:"name"`
	IsOrganizer bool                   `json:"is_organizer"`
	ImageURL    string                 `json:"image_url"`
	Message     string                 `json:"message"`
	Skills      map[string]string      `json:"skills"`
	URLs        map[entity.URLs]string `json:"urls"`
}
