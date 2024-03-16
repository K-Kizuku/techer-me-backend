//go:build wireinject
// +build wireinject

package di

import (
	"github.com/K-Kizuku/techer-me-backend/internal/app/handler"
	evh "github.com/K-Kizuku/techer-me-backend/internal/app/handler/event"
	eh "github.com/K-Kizuku/techer-me-backend/internal/app/handler/exchange"
	ih "github.com/K-Kizuku/techer-me-backend/internal/app/handler/image"
	uh "github.com/K-Kizuku/techer-me-backend/internal/app/handler/user"
	evr "github.com/K-Kizuku/techer-me-backend/internal/app/repository/event"
	er "github.com/K-Kizuku/techer-me-backend/internal/app/repository/exchange"
	ir "github.com/K-Kizuku/techer-me-backend/internal/app/repository/image"
	ur "github.com/K-Kizuku/techer-me-backend/internal/app/repository/user"
	evs "github.com/K-Kizuku/techer-me-backend/internal/app/service/event"
	es "github.com/K-Kizuku/techer-me-backend/internal/app/service/exchange"
	is "github.com/K-Kizuku/techer-me-backend/internal/app/service/image"
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
		ir.New,
		ih.New,
		is.New,
		handler.New,
	)
	return &handler.Root{}
}
