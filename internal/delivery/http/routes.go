package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/timmbarton/example/internal/usecase"
	"github.com/timmbarton/example/pkg/errlist"
	"github.com/timmbarton/response"
	"github.com/timmbarton/utils/tracing"
)

func (h *handler) routeSomeSomeGet(c *fiber.Ctx) error {
	ctx, span := tracing.NewSpan(c.UserContext())
	defer span.End()
	c.SetUserContext(ctx)

	req := usecase.FooBarRequest{}

	err := c.QueryParser(&req)
	if err != nil {
		return errlist.ErrInternal
	}

	data, err := h.uc.FooBar(ctx, req)
	if err != nil {
		return err
	}

	return response.OkWithData(c, data)
}
