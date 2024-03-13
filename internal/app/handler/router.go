package handler

import "github.com/K-Kizuku/techer-me-backend/internal/app/handler/user"

type Root struct {
	UserHandler *user.Handler
}

func New(userHandler *user.Handler) *Root {
	return &Root{
		UserHandler: userHandler,
	}
}
