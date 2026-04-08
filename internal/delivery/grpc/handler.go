package grpc

import (
	"context"

	"github.com/timmbarton/example/internal/usecase"
	proto "github.com/timmbarton/example/pkg/api/grpc"
	"github.com/timmbarton/utils/tracing"
)

type handler struct {
	proto.UnimplementedExampleServiceServer
	uc UseCase
}

func (h *handler) FooBar(ctx context.Context, req *proto.FooBarRequest) (res *proto.FooBarResponse, err error) {
	ctx, span := tracing.NewSpan(ctx)
	defer span.End()

	res = new(proto.FooBarResponse)

	res.Data, err = h.uc.FooBar(ctx, usecase.FooBarRequest{
		Data: req.Data,
	})
	if err != nil {
		return res, err
	}

	return res, nil
}
