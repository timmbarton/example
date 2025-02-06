package usecase

import (
	"github.com/timmbarton/example/internal/cache"
	"github.com/timmbarton/example/internal/repository"
)

type UseCases struct {
	Webmaster WebmasterUseCase
	Admin     AdminUseCase
}

type Config struct {
	Webmaster WebmasterConfig
}

func New(cfg Config, r *repository.Repositories, c cache.Cache) *UseCases {
	return &UseCases{
		Webmaster: &webmasterUseCase{
			cfg: cfg.Webmaster,
			r:   r,
			c:   c,
		},
	}
}

type WebmasterUseCase interface {
}

type AdminUseCase interface {
}
