package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/timmbarton/layout/components/httpserver"
	"github.com/timmbarton/layout/lifecycle"

	"github.com/timmbarton/example/internal/usecase"
)

func New(cfg httpserver.Config, uc *usecase.UseCases) lifecycle.Lifecycle {
	s := &httpserver.DefaultServer{}

	h := &handler{
		uc: uc,
		v:  validator.New(),
	}

	s.Init(cfg, h.bind)

	return s
}
