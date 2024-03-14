package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/K-Kizuku/techer-me-backend/pkg/firebase"
)

type Key string

const UserIDKey Key = "userID"

func FirebaseAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		app, err := firebase.InitFirebaseApp()
		if err != nil {
			http.Error(w, "Firebase initialization failed", http.StatusInternalServerError)
			slog.Error("Firebase initialization failed", "error", err.Error())
			return
		}

		client, err := app.Auth(ctx)
		if err != nil {
			http.Error(w, "Failed to initialize Auth client", http.StatusInternalServerError)
			slog.Error("Failed to initialize Auth client", "error", err.Error())
			return
		}

		authHeader := r.Header.Get("Authorization")
		idToken := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify ID token
		token, err := client.VerifyIDToken(ctx, idToken)
		if err != nil {
			http.Error(w, "Invalid or expired ID token", http.StatusUnauthorized)
			slog.Error("Invalid or expired ID token", "error", err.Error())
			return
		}

		// Token is valid, attach user information to the context
		ctx = context.WithValue(ctx, UserIDKey, token.UID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
