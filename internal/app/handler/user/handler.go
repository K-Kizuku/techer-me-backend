package user

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/K-Kizuku/techer-me-backend/internal/app/handler/schema"
	"github.com/K-Kizuku/techer-me-backend/internal/app/service/user"
	"github.com/K-Kizuku/techer-me-backend/pkg/errors"
	"github.com/K-Kizuku/techer-me-backend/pkg/middleware"
)

type Handler struct {
	userService user.IUserService
}

func New(userService user.IUserService) *Handler {
	return &Handler{
		userService: userService,
	}
}

func (h *Handler) CreateUserByFirebaseID() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req schema.CreateUserInput
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return errors.New(http.StatusBadRequest, err)
		}
		if err := h.userService.CreateUserByFirebaseID(r.Context(), req.UserID); err != nil {
			return err
		}
		if err := h.userService.CreateUserDetailByFirebaseID(r.Context(), req); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return err
		}
		w.WriteHeader(http.StatusCreated)

		fmt.Fprint(w, "OK")
		return nil
	}
}

// @Summary 自分の情報
// @Description 自分の情報を取得する
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} schema.GetMeOutput "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Security Bearer
// @Router /me [get]
func (h *Handler) GetMe() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {

		userID := r.Context().Value(middleware.UserIDKey).(string)

		user, err := h.userService.GetByID(r.Context(), userID)
		if err != nil {
			return err
		}
		var events []schema.Event
		for _, event := range user.Events {
			events = append(events, schema.Event{
				EventID:    event.ID,
				Name:       event.Name,
				OwnerID:    event.OwnerID,
				StartedAt:  event.StartedAt,
				FinishedAt: event.FinishedAt,
				Message:    event.Message,
				ImageURL:   event.ImageURL,
			})
		}

		res := schema.GetMeOutput{
			UserID:      user.ID,
			Name:        user.Name,
			IsOrganizer: user.IsOrganizer,
			ImageURL:    user.ImageURL,
			Message:     user.Message,
			Skills:      user.Skills,
			URLs:        user.URLs,
			Events:      events,
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			return errors.New(http.StatusInternalServerError, err)
		}
		return nil
	}
}

// @Summary ユーザー情報取得
// @Description 指定したユーザーの情報を取得する
// @Tags User
// @Accept json
// @Produce json
// @Param id body schema.GetByIDInput true "ユーザーID"
// @Success 200 {object} schema.GetByIDOutput "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Security Bearer
// @Router /user [get]
func (h *Handler) GetByID() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req schema.GetByIDInput
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return errors.New(http.StatusBadRequest, err)
		}
		user, err := h.userService.GetByID(r.Context(), req.UserID)
		if err != nil {
			return err
		}
		var events []schema.Event
		for _, event := range user.Events {
			events = append(events, schema.Event{
				EventID:    event.ID,
				Name:       event.Name,
				OwnerID:    event.OwnerID,
				StartedAt:  event.StartedAt,
				FinishedAt: event.FinishedAt,
				Message:    event.Message,
				ImageURL:   event.ImageURL,
			})
		}
		res := schema.GetByIDOutput{
			UserID:      user.ID,
			Name:        user.Name,
			IsOrganizer: user.IsOrganizer,
			ImageURL:    user.ImageURL,
			Message:     user.Message,
			Skills:      user.Skills,
			URLs:        user.URLs,
			Events:      events,
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			return errors.New(http.StatusInternalServerError, err)
		}
		return nil
	}

}

func (h *Handler) GetEventByID() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {

		userID := r.Context().Value(middleware.UserIDKey).(string)

		events, err := h.userService.GetEventByID(r.Context(), userID)
		if err != nil {
			return err
		}
		var res schema.GetEventByIDOutput
		res.Events = make([]schema.Event, 0)
		for _, event := range events {
			res.Events = append(res.Events, schema.Event{
				EventID:    event.ID,
				Name:       event.Name,
				OwnerID:    event.OwnerID,
				StartedAt:  event.StartedAt,
				FinishedAt: event.FinishedAt,
				Message:    event.Message,
				ImageURL:   event.ImageURL,
			})
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			return errors.New(http.StatusInternalServerError, err)
		}
		return nil
	}
}
