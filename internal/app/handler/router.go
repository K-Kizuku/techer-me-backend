package handler

import (
	"github.com/K-Kizuku/techer-me-backend/internal/app/handler/event"
	"github.com/K-Kizuku/techer-me-backend/internal/app/handler/exchange"
	"github.com/K-Kizuku/techer-me-backend/internal/app/handler/user"
)

type Root struct {
	UserHandler     *user.Handler
	ExchangeHandler *exchange.Handler
	EventHandler    *event.Handler
}

func New(userHandler *user.Handler, exchangeHandler *exchange.Handler, eventHandler *event.Handler) *Root {
	return &Root{
		UserHandler:     userHandler,
		ExchangeHandler: exchangeHandler,
		EventHandler:    eventHandler,
	}
}
