package handler

import (
	"log/slog"
	"net/http"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		slog.Error("error handling request", "error", err)
	}
}
