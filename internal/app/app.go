package app

import (
	"github.com/timmbarton/layout/components/postgresconn"
	"github.com/timmbarton/layout/components/redisconn"
	"github.com/timmbarton/layout/components/signoz"
	"github.com/timmbarton/layout/executor"
	"github.com/timmbarton/layout/template"

	"github.com/timmbarton/example/internal/cache"
	"github.com/timmbarton/example/internal/config"
	"github.com/timmbarton/example/internal/delivery/grpc"
	"github.com/timmbarton/example/internal/delivery/http"
	"github.com/timmbarton/example/internal/repository"
	"github.com/timmbarton/example/internal/usecase"
)

func New(cfg config.Config) (executor.App, error) {
	a := &template.App{}

	sn, err := signoz.New(cfg.SigNoz)

	pg, err := postgresconn.New(cfg.Postgres)
	if err != nil {
		return nil, err
	}

	rd, err := redisconn.New(cfg.Redis)
	if err != nil {
		return nil, err
	}

	r := repository.New(pg.DB())
	c := cache.NewCache(rd)
	uc := usecase.New(cfg.UseCase, r, c)
	httpServer := http.New(cfg.HTTP, uc)

	grpcServer, err := grpc.New(cfg.GRPC, uc)
	if err != nil {
		return nil, err
	}

	a.AddComponents(
		sn,
		pg,
		rd,
		httpServer,
		grpcServer,
	)

	return a, nil
}
