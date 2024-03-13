//go:build wireinject
// +build wireinject

package di

import (
	"github.com/K-Kizuku/techer-me-backend/internal/app/handler"
	uh "github.com/K-Kizuku/techer-me-backend/internal/app/handler/user"
	ur "github.com/K-Kizuku/techer-me-backend/internal/app/repository/user"
	us "github.com/K-Kizuku/techer-me-backend/internal/app/service/user"
	"github.com/K-Kizuku/techer-me-backend/pkg/db"
	"github.com/google/wire"
)

func InitHandler() *handler.Root {
	wire.Build(
		db.New,
		ur.New,
		uh.New,
		us.New,
		handler.New,
	)
	return &handler.Root{}
}
