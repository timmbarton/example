package grpc

import (
	"github.com/timmbarton/layout/components/grpcserver"
	"github.com/timmbarton/layout/lifecycle"

	"github.com/timmbarton/example/internal/usecase"
)

func New(cfg grpcserver.DefaultServerConfig, uc *usecase.UseCases) (lifecycle.Lifecycle, error) {
	s := &grpcserver.DefaultServer{}

	s.Init(cfg)

	//proto.Register***Server(s.grpcServer, &handler{uc: uc})

	return s, nil
}
