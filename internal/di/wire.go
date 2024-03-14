//go:build wireinject
// +build wireinject

package di

import (
	"github.com/K-Kizuku/techer-me-backend/internal/app/handler"
	evh "github.com/K-Kizuku/techer-me-backend/internal/app/handler/event"
	eh "github.com/K-Kizuku/techer-me-backend/internal/app/handler/exchange"
	uh "github.com/K-Kizuku/techer-me-backend/internal/app/handler/user"
	evr "github.com/K-Kizuku/techer-me-backend/internal/app/repository/event"
	er "github.com/K-Kizuku/techer-me-backend/internal/app/repository/exchange"
	ur "github.com/K-Kizuku/techer-me-backend/internal/app/repository/user"
	evs "github.com/K-Kizuku/techer-me-backend/internal/app/service/event"
	es "github.com/K-Kizuku/techer-me-backend/internal/app/service/exchange"
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
		er.New,
		eh.New,
		es.New,
		evr.New,
		evh.New,
		evs.New,
		handler.New,
	)
	return &handler.Root{}
}
