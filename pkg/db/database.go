package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/K-Kizuku/techer-me-backend/pkg/config"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New() *sqlx.DB {
	db, err := Init()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Init() (*sqlx.DB, error) {
	if config.Mode == "dev" {
		db, err := sqlx.Open("mysql", config.DB_DSN)
		if err != nil {
			return nil, err
		}
		return db, nil
	} else if config.Mode == "prod" {
		d, err := cloudsqlconn.NewDialer(context.Background())
		if err != nil {
			return nil, err
		}

		var opts []cloudsqlconn.DialOption

		mysql.RegisterDialContext("cloudsqlconn",
			func(ctx context.Context, addr string) (net.Conn, error) {
				return d.Dial(ctx, config.InstanceConnectionName, opts...)
			},
		)

		dbURI := fmt.Sprintf(
			"%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true",
			config.DBUser, config.DBPassword, config.DBName,
		)

		db, err := sqlx.Open("mysql", dbURI)
		if err != nil {
			return nil, err
		}

		return db, err
	}
	return nil, errors.New("mode is invalid")
}
