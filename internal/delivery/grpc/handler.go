package grpc

import (
	"github.com/timmbarton/example/internal/usecase"
)

type handler struct {
	//proto.Unimplemented...
	uc *usecase.UseCases
}
