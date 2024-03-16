package event

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/K-Kizuku/techer-me-backend/internal/app/handler/schema"
	"github.com/K-Kizuku/techer-me-backend/internal/app/service/event"
	"github.com/K-Kizuku/techer-me-backend/pkg/errors"
	"github.com/K-Kizuku/techer-me-backend/pkg/middleware"
)

type Handler struct {
	eventService event.IEventService
}

func New(eventService event.IEventService) *Handler {
	return &Handler{
		eventService: eventService,
	}
}

// @Summary Join an event
// @Description Join an event by event ID
// @Tags Event
// @Accept json
// @Produce json
// @Param event_id query string true "Event ID"
// @Success 201 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /event/join [post]
func (h *Handler) Join() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		userID := r.Context().Value(middleware.UserIDKey).(string)
		eventID := r.PathValue("event_id")
		if eventID == "" {
			return errors.New(http.StatusBadRequest, fmt.Errorf("event_id is required"))
		}
		if err := h.eventService.Join(r.Context(), eventID, userID); err != nil {
			return err
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "OK")
		return nil
	}
}

// @Summary イベント作成
// @Description イベント作成のためのエンドポイント
// @Tags Event
// @Accept json
// @Produce json
// @Param event body schema.CreateEventInput true "Event request body"
// @Success 201 {object} schema.CreateEventOutput "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /events [post]
func (h *Handler) Create() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req schema.CreateEventInput
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return errors.New(http.StatusBadRequest, err)
		}
		eventID, err := h.eventService.Create(r.Context(), &req)
		if err != nil {
			return err
		}
		e := schema.CreateEventOutput{
			EventID: eventID,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(e); err != nil {
			return errors.New(http.StatusInternalServerError, err)
		}
		return nil
	}
}

// @Summary イベント詳細の取得
// @Description イベント詳細を取得する
// @Tags Event
// @Accept json
// @Produce json
// @Param event_id path string true "Event ID"
// @Success 200 {object} schema.GetEventDetailByIDOutput
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /events/{event_id} [get]
func (h *Handler) GetEventDetailByID() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		eventID := r.PathValue("event_id")
		if eventID == "" {
			return errors.New(http.StatusBadRequest, fmt.Errorf("event_id is required"))
		}
		event, err := h.eventService.SelectByID(r.Context(), eventID)
		if err != nil {
			return err
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(event); err != nil {
			return errors.New(http.StatusInternalServerError, err)
		}
		return nil
	}
}
