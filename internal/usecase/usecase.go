package usecase

import (
	"context"

	"github.com/timmbarton/utils/tracing"
)

func New(cfg Config, r Repository, c Cache) *UseCase {
	return &UseCase{
		cfg: cfg,
		r:   r,
		c:   c,
	}
}

type Config struct {
}

type UseCase struct {
	cfg Config
	r   Repository
	c   Cache
}

type FooBarRequest struct {
	Data string `json:"data" validate:"required,max=100"`
}

func (uc *UseCase) FooBar(ctx context.Context, req FooBarRequest) (res []int64, err error) {
	ctx, span := tracing.NewSpan(ctx)
	defer span.End()

	// some work
	if req.Data == "empty" {
		return res, err
	}

	res, err = uc.r.GetFooBar(ctx)
	if err != nil {
		return res, err
	}

	return nil, nil
}
