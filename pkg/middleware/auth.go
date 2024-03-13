package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/K-Kizuku/techer-me-backend/pkg/firebase"
)

func FirebaseAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app, err := firebase.InitFirebaseApp()
		if err != nil {
			http.Error(w, "Firebase initialization failed", http.StatusInternalServerError)
			return
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			http.Error(w, "Failed to initialize Auth client", http.StatusInternalServerError)
			return
		}

		authHeader := r.Header.Get("Authorization")
		idToken := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify ID token
		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			http.Error(w, "Invalid or expired ID token", http.StatusUnauthorized)
			return
		}

		// Token is valid, attach user information to the context
		ctx := context.WithValue(r.Context(), "uid", token.UID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
