package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/K-Kizuku/techer-me-backend/internal/di"
	"github.com/K-Kizuku/techer-me-backend/pkg/config"
	"github.com/K-Kizuku/techer-me-backend/pkg/handler"
	"github.com/K-Kizuku/techer-me-backend/pkg/log"
	"github.com/K-Kizuku/techer-me-backend/pkg/middleware"
	"github.com/rs/cors"
)

func main() {
	config.LoadEnv()

	h := di.InitHandler()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World")
	})

	// mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
	// 	err := db.Ping()
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		fmt.Fprint(w, err.Error())
	// 	}
	// 	w.WriteHeader(http.StatusOK)
	// 	fmt.Fprint(w, "pong")
	// })

	mux.Handle("POST /users", handler.AppHandler(h.UserHandler.CreateUserByFirebaseID()))

	mux.Handle("GET /me", middleware.FirebaseAuth(handler.AppHandler(h.UserHandler.GetMe())))

	c := cors.AllowAll()
	handler := middleware.Chain(mux, middleware.Context, c.Handler, middleware.Recover, middleware.Logger)

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	go func() {
		log.Start()
		if err := server.ListenAndServe(); err != nil {
			slog.Error("server error", "error", err.Error())
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server error", "error", err.Error())
	}
}
