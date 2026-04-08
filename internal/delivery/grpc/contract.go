package grpc

import (
	"context"

	"github.com/timmbarton/example/internal/usecase"
)

type UseCase interface {
	FooBar(ctx context.Context, req usecase.FooBarRequest) (data []int64, err error)
}
