package exchange

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/K-Kizuku/techer-me-backend/internal/app/handler/schema"
	"github.com/K-Kizuku/techer-me-backend/internal/app/service/exchange"
	"github.com/K-Kizuku/techer-me-backend/pkg/errors"
)

type Handler struct {
	exchangeService exchange.IExchangeService
}

func New(exchengeService exchange.IExchangeService) *Handler {
	return &Handler{
		exchangeService: exchengeService,
	}
}

// @Summary Create Exchange
// @Description 名刺交換のためのエンドポイント
// @Tags Exchange
// @Accept json
// @Produce json
// @Param exchange body schema.CreateExchangeInput true "Exchange request body"
// @Success 201 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Security Bearer
// @Router /exchanges [post]
func (h *Handler) CreateExchange() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req schema.CreateExchangeInput
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return errors.New(http.StatusBadRequest, err)
		}
		if err := h.exchangeService.Create(r.Context(), &req); err != nil {
			return err
		}
		w.WriteHeader(http.StatusCreated)

		fmt.Fprint(w, "OK")
		return nil
	}
}
