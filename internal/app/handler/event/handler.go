package event

import (
	"fmt"
	"net/http"

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
