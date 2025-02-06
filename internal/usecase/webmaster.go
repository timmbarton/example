package usecase

import (
	"github.com/timmbarton/example/internal/cache"
	"github.com/timmbarton/example/internal/repository"
)

type WebmasterConfig struct {
}

type webmasterUseCase struct {
	cfg WebmasterConfig
	r   *repository.Repositories
	c   cache.Cache
}
