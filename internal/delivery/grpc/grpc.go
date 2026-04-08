package grpc

import (
	proto "github.com/timmbarton/example/pkg/api/grpc"
	"github.com/timmbarton/layout/components/grpcserver"
	"github.com/timmbarton/layout/lifecycle"
)

func New(cfg grpcserver.DefaultServerConfig, uc UseCase) (lifecycle.Lifecycle, error) {
	s := &grpcserver.DefaultServer{}

	err := s.Init(cfg)
	if err != nil {
		return nil, err
	}

	s.RegisterService(&proto.ExampleService_ServiceDesc, &handler{uc: uc})

	return s, nil
}
