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

	_ "github.com/K-Kizuku/techer-me-backend/docs"
	"github.com/K-Kizuku/techer-me-backend/internal/di"
	"github.com/K-Kizuku/techer-me-backend/pkg/config"
	"github.com/K-Kizuku/techer-me-backend/pkg/handler"
	"github.com/K-Kizuku/techer-me-backend/pkg/log"
	"github.com/K-Kizuku/techer-me-backend/pkg/middleware"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host server-u7kyixk36q-an.a.run.app
// @BasePath /
// @SecurityDefinitions.apiKey Bearer
// @in header
// @name Authorization
func main() {
	config.LoadEnv()

	h := di.InitHandler()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World")
	})

	mux.HandleFunc("GET /swagger/*", httpSwagger.Handler(
		httpSwagger.URL("https://server-u7kyixk36q-an.a.run.app/swagger/doc.json"), //The url pointing to API definition
	))

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
	mux.Handle("GET /users", middleware.FirebaseAuth(handler.AppHandler(h.UserHandler.GetByID())))
	mux.Handle("PUT /me", middleware.FirebaseAuth(handler.AppHandler(h.UserHandler.Update())))

	mux.Handle("POST /exchanges", middleware.FirebaseAuth(handler.AppHandler(h.ExchangeHandler.CreateExchange())))
	mux.Handle("GET /exchanges", middleware.FirebaseAuth(handler.AppHandler(h.ExchangeHandler.GetExchanges())))

	mux.Handle("POST /events/join/{event_id}", middleware.FirebaseAuth(handler.AppHandler(h.EventHandler.Join())))

	c := cors.AllowAll()
	handler := middleware.Chain(mux, middleware.Context, c.Handler, middleware.Recover, middleware.Logger)

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	go func() {
		log.Start()
		if err := server.ListenAndServe(); err != nil {
			slog.Error("server error", "error", err)
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
