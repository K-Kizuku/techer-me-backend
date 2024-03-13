package user

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/K-Kizuku/techer-me-backend/internal/app/handler/schema"
	"github.com/K-Kizuku/techer-me-backend/internal/app/service/user"
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
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return err
		}
		if err := h.userService.CreateUserByFirebaseID(r.Context(), req.UserID); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return err
		}
		w.WriteHeader(http.StatusCreated)

		fmt.Fprint(w, "OK")
		return nil
	}
}

func (h *Handler) GetMe() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {

		userID := r.Context().Value(middleware.UserIDKey).(string)
		fmt.Println(userID)
		fmt.Println("fuga")

		user, err := h.userService.GetByID(r.Context(), userID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return err
		}
		res := schema.GetMeOutput{
			UserID:      user.ID,
			Name:        user.Name,
			IsOrganizer: user.IsOrganizer,
			ImageURL:    user.ImageURL,
			Message:     user.Message,
			Skills:      user.Skills,
			URLs:        user.URLs,
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return err
		}
		return nil
	}
}
