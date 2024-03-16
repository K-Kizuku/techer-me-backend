package image

import (
	"encoding/json"
	"net/http"

	"github.com/K-Kizuku/techer-me-backend/internal/app/handler/schema"
	"github.com/K-Kizuku/techer-me-backend/internal/app/service/image"
	"github.com/K-Kizuku/techer-me-backend/pkg/errors"
)

type Handler struct {
	imageService image.IImageService
}

func New(imageService image.IImageService) *Handler {
	return &Handler{
		imageService: imageService,
	}
}

// @Summary Generate signed URL
// @Description Generate signed URL for uploading image
// @Tags Image
// @Accept json
// @Produce json
// @Param generateImage body schema.GenerateImageInput true "Generate image request body"
// @Success 200 {object} schema.GenerateImageOutput "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Security Bearer
// @Router /images/signed-url [post]
func (h *Handler) GenerateSignedURL() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req schema.GenerateImageInput
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return errors.New(http.StatusBadRequest, err)
		}
		url, err := h.imageService.GenerateSignedURL(r.Context(), &req)
		if err != nil {
			return err
		}
		res := schema.GenerateImageOutput{
			URL: url.URL,
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			return errors.New(http.StatusInternalServerError, err)
		}
		return nil
	}
}
