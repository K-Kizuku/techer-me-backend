package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"github.com/K-Kizuku/techer-me-backend/pkg/config"
	"google.golang.org/api/option"
)

func InitFirebaseApp() (*firebase.App, error) {
	opt := option.WithCredentialsFile(config.GoogleApplicationCredentials)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	return app, nil
}
