package image

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/entity"
	"github.com/K-Kizuku/techer-me-backend/internal/domain/repository/image"
	"github.com/K-Kizuku/techer-me-backend/pkg/config"
	"github.com/K-Kizuku/techer-me-backend/pkg/errors"
	"github.com/jmoiron/sqlx"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iam/v1"
)

type repository struct {
}

func New(conn *sqlx.DB) image.IRepository {
	return &repository{}
}

func (r *repository) GenerateSignedURL(ctx context.Context, image entity.Image) (string, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", errors.New(http.StatusInternalServerError, err)
	}
	defer client.Close()
	cred, err := google.DefaultClient(ctx, iam.CloudPlatformScope)
	if err != nil {
		return "", errors.New(http.StatusInternalServerError, err)
	}
	iamService, err := iam.New(cred)
	opts := &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		GoogleAccessID: config.GoogleAccessID,
		Method:         "PUT",
		Expires:        time.Now().Add(15 * time.Minute),
		ContentType:    "image/" + image.Type,
		SignBytes: func(b []byte) ([]byte, error) {
			resp, err := iamService.Projects.ServiceAccounts.SignBlob(
				config.Name, &iam.SignBlobRequest{
					BytesToSign: base64.StdEncoding.EncodeToString(b)},
			).Context(ctx).Do()
			if err != nil {
				return nil, errors.New(http.StatusInternalServerError, err)
			}
			return base64.StdEncoding.DecodeString(resp.Signature)
		},
	}
	bucket := config.Basket
	object := image.ObjectName
	signedURL, err := storage.SignedURL(bucket, object, opts)
	if err != nil {
		return "", errors.New(http.StatusInternalServerError, err)
	}
	fmt.Println(signedURL)

	return signedURL, nil
}
